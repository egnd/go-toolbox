package graceful_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/graceful"

	"github.com/stretchr/testify/assert"
)

func Test_graceful(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	_, cancelFunc := graceful.Init(context.TODO())
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancelFunc()
	}()

	graceful.Register(func() error {
		wg.Wait()
		return nil
	}, func() error {
		wg.Done()
		return errors.New("some error")
	})

	assert.Errorf(t, graceful.Wait(), "asf")
}

func Test_graceful_register_errors(t *testing.T) {
	cases := []struct {
		callbacks []graceful.Callback
		err       error
	}{
		{
			err: graceful.ErrCallbackCnt,
		},
		{
			callbacks: []graceful.Callback{func() error { return nil }, func() error { return nil }, func() error { return nil }},
			err:       graceful.ErrCallbackCnt,
		},
		{
			callbacks: []graceful.Callback{nil},
			err:       graceful.ErrStartFunc,
		},
		{
			callbacks: []graceful.Callback{func() error { return nil }, nil},
			err:       graceful.ErrStopFunc,
		},
		{
			callbacks: []graceful.Callback{func() error { return nil }},
		},
		{
			callbacks: []graceful.Callback{func() error { return nil }, func() error { return nil }},
		},
	}
	graceful.Init(context.TODO())
	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			if test.err == nil {
				assert.NoError(t, graceful.Register(test.callbacks...))
			} else {
				assert.ErrorIs(t, graceful.Register(test.callbacks...), test.err)
			}
		})
	}
}
