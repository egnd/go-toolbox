package graceful

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"
)

var (
	group     *errgroup.Group //nolint:gochecknoglobals
	gctx      context.Context //nolint:gochecknoglobals
	callbacks []Callback      //nolint:gochecknoglobals
)

// Init subscribes to os signals and registers cancel func and group context.
func Init(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	var cancelFunc context.CancelFunc

	if len(signals) > 0 {
		ctx, cancelFunc = signal.NotifyContext(ctx, signals...)
	} else {
		ctx, cancelFunc = context.WithCancel(ctx)
	}

	group, gctx = errgroup.WithContext(ctx)

	return ctx, cancelFunc
}

// Register registers some service to listen and its "stop" function.
func Register(funcs ...Callback) error {
	if len(funcs) == 0 || len(funcs) > 2 {
		return ErrCallbackCnt
	}

	if funcs[0] == nil {
		return ErrStartFunc
	}

	group.Go(funcs[0])

	if len(funcs) == 1 {
		return nil
	}

	if funcs[1] == nil {
		return ErrStopFunc
	}

	callbacks = append(callbacks, funcs[1])

	return nil
}

// Wait registers executor for "stop" callbacks and blocks until all registered callbacks will end.
func Wait() error {
	err := Register(func() (err error) { //nolint:nonamedreturns
		<-gctx.Done()

		for _, callback := range callbacks {
			if cErr := callback(); cErr != nil {
				err = multierror.Append(err, cErr)
			}
		}

		return
	})
	if err != nil {
		return fmt.Errorf("graceful wait register error: %w", err)
	}

	if err = group.Wait(); err != nil {
		err = fmt.Errorf("graceful wait error: %w", err)
	}

	return err
}
