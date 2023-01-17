package pools_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/pipelines/mocks"
	"github.com/egnd/go-toolbox/pipelines/pools"
	"github.com/stretchr/testify/assert"
)

func newTask(id string) *mocks.Task {
	task := &mocks.Task{}

	task.On("ID").Return(id)
	task.On("Do").Return(nil)

	return task
}

func Test_Delayer(t *testing.T) {
	cases := []struct {
		ttl      time.Duration
		pausesMs []int
	}{
		{
			ttl:      200 * time.Millisecond,
			pausesMs: []int{100, 100, 300},
		},
	}
	for k, test := range cases {
		t.Run(fmt.Sprint(k), func(tt *testing.T) {
			task := newTask("some-id")
			defer task.AssertExpectations(tt)

			pool := pools.NewDelayer(context.TODO(), test.ttl)

			for _, durMs := range test.pausesMs {
				assert.NoError(tt, pool.Push(task))
				time.Sleep(time.Millisecond * time.Duration(durMs))
			}

			pool.Wait()
			assert.NoError(tt, pool.Close())
		})
	}
}

func Test_Delayer_Close(t *testing.T) {
	pool := pools.NewDelayer(context.TODO(), time.Minute)

	assert.NoError(t, pool.Push(newTask("some-id")))

	go func() {
		time.Sleep(200 * time.Millisecond)
		assert.NoError(t, pool.Close())
	}()

	pool.Wait()
}

func Test_Delayer_Ctx_Cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())

	pool := pools.NewDelayer(ctx, time.Minute)

	assert.NoError(t, pool.Push(newTask("some-id")))

	go func() {
		time.Sleep(200 * time.Millisecond)
		cancel()
	}()

	pool.Wait()
}

func Test_Delayer_EmptyTask(t *testing.T) {
	pool := pools.NewDelayer(context.TODO(), time.Minute)
	assert.NoError(t, pool.Push(nil))
	pool.Wait()
}
