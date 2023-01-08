package graceful_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/graceful"
	"github.com/hashicorp/go-multierror"

	"github.com/stretchr/testify/assert"
)

func Test_Listener(t *testing.T) {
	cases := []struct {
		signals   []os.Signal
		callbacks func(context.Context, *graceful.Listener)
		err       error
	}{
		{
			signals: []os.Signal{os.Interrupt},
			callbacks: func(ctx context.Context, l *graceful.Listener) {
				l.Register(func() { <-ctx.Done() }, func() { <-ctx.Done() })
			},
		},
		{},
		{
			callbacks: func(ctx context.Context, l *graceful.Listener) {
				l.Register(func() bool { <-ctx.Done(); return true })
			},
			err: multierror.Append(nil, errors.New("invalid start callback type - func() bool")),
		},
		{
			callbacks: func(ctx context.Context, l *graceful.Listener) {
				l.Register(func() { <-ctx.Done() }, func() error { <-ctx.Done(); return errors.New("some error") })
			},
			err: multierror.Append(nil, errors.New("some error")),
		},
		{
			callbacks: func(ctx context.Context, l *graceful.Listener) {
				l.Register(func() { <-ctx.Done() }, func() bool { <-ctx.Done(); return true })
			},
			err: multierror.Append(nil, errors.New("invalid stop callback type - func() bool")),
		},
		{
			callbacks: func(ctx context.Context, l *graceful.Listener) {
				l.Register()
			},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			listener, ctx, cancelFunc := graceful.NewListener(context.TODO(), test.signals...)

			go func() {
				time.Sleep(100 * time.Millisecond)
				cancelFunc()
			}()

			if test.callbacks != nil {
				test.callbacks(ctx, listener)
			}

			if test.err == nil {
				assert.NoError(t, listener.Wait())
			} else {
				assert.EqualError(t, listener.Wait(), test.err.Error())
			}
		})
	}
}
