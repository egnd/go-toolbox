package graceful

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"
)

// Listener listens errorgroup goroutines.
type Listener struct {
	group *errgroup.Group
	gctx  context.Context //nolint:containedctx
	start []interface{}
	stop  []interface{}
}

// NewListener create listener for os signals and registers cancel func and group context.
func NewListener(ctx context.Context, signals ...os.Signal) (*Listener, context.Context, context.CancelFunc) {
	var cancelFunc context.CancelFunc

	if len(signals) > 0 {
		ctx, cancelFunc = signal.NotifyContext(ctx, signals...)
	} else {
		ctx, cancelFunc = context.WithCancel(ctx)
	}

	group, gctx := errgroup.WithContext(ctx)

	return &Listener{ //nolint:exhaustruct
		group: group,
		gctx:  gctx,
	}, ctx, cancelFunc
}

// Register registers some service to listen and its "stop" function.
func (l *Listener) Register(funcs ...interface{}) {
	if len(funcs) == 0 {
		return
	}

	l.start = append(l.start, funcs[0])

	if len(funcs) == 1 {
		return
	}

	l.stop = append(l.stop, funcs[1])
}

func (l *Listener) startCallbacks() (cnt int, err error) {
	for _, callback := range l.start {
		switch startFunc := callback.(type) {
		case Callback:
			cnt++

			l.group.Go(func() error {
				startFunc()

				return nil
			})
		case CallbackWithError:
			cnt++

			l.group.Go(startFunc)
		case nil:
		default:
			err = multierror.Append(err, fmt.Errorf("invalid start callback type - %T", callback))
		}
	}

	return
}

func (l *Listener) stopCallbacks() (err error) {
	for _, callback := range l.stop {
		switch stopFunc := callback.(type) {
		case Callback:
			stopFunc()
		case CallbackWithError:
			if cErr := stopFunc(); cErr != nil {
				err = multierror.Append(err, cErr)
			}
		case nil:
		default:
			err = multierror.Append(err, fmt.Errorf("invalid stop callback type - %T", callback))
		}
	}

	return err
}

// Wait registers executor for "stop" callbacks and blocks until all registered callbacks will end.
func (l *Listener) Wait() error {
	l.Register(func() error {
		<-l.gctx.Done()

		return l.stopCallbacks()
	})

	if total, err := l.startCallbacks(); err != nil || total < 2 {
		return err
	}

	return l.group.Wait()
}
