package pools_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/pipelines/assign"
	"github.com/egnd/go-toolbox/pipelines/mocks"
	"github.com/egnd/go-toolbox/pipelines/pools"

	"github.com/stretchr/testify/assert"
)

func Test_HashPool(t *testing.T) {
	cases := []struct {
		threadsCnt int
		tasksCnt   int
	}{
		{3, 100},
		{1, 20},
		{10, 200},
	}
	for k, test := range cases {
		t.Run(fmt.Sprint(k), func(tt *testing.T) {
			task := &mocks.Task{}
			defer task.AssertExpectations(tt)
			task.On("ID").Return("some-id").Times(test.tasksCnt)
			task.On("Do").After(time.Duration(rand.Intn(10)) * time.Millisecond).Times(test.tasksCnt).Return(nil)

			pipe := pools.NewHashPool(test.threadsCnt, 0, nil, assign.Sticky)
			for i := 0; i < test.tasksCnt; i++ {
				assert.NoError(tt, pipe.Push(task))
			}
			pipe.Wait()
			assert.NoError(tt, pipe.Close())
		})
	}
}

func Test_HashPool_Errors(t *testing.T) {
	pool := pools.NewHashPool(1, 0, nil, assign.Sticky)

	assert.NoError(t, pool.Close())
	assert.EqualError(t, pool.Close(), "pool close err: close of closed channel")
	assert.EqualError(t, pool.Push(nil), "pool push err: send on closed channel")
}

func Test_HashPool_NoThreads(t *testing.T) {
	assert.PanicsWithValue(t, "HashPool requires at least 1 thread", func() {
		pools.NewHashPool(0, 0, nil, assign.Sticky)
	})
}
