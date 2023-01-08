package graceful

import (
	"context"
	"os"
)

var defaultListener *Listener //nolint:gochecknoglobals

// Init subscribes to os signals and registers cancel func and group context.
func Init(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	var cancelFunc context.CancelFunc

	defaultListener, ctx, cancelFunc = NewListener(ctx, signals...)

	return ctx, cancelFunc
}

// Register registers some service to listen and its "stop" function.
func Register(funcs ...interface{}) error {
	if defaultListener == nil {
		return ErrEmptyListener
	}

	defaultListener.Register(funcs...)

	return nil
}

// Wait registers executor for "stop" callbacks and blocks until all registered callbacks will end.
func Wait() error {
	if defaultListener == nil {
		return ErrEmptyListener
	}

	return defaultListener.Wait()
}
