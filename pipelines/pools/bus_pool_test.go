package pools_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/pipelines/mocks"
	"github.com/egnd/go-toolbox/pipelines/pools"

	"github.com/stretchr/testify/assert"
)

func Test_PipelinePool(t *testing.T) {
	cases := []struct {
		tasksCnt   int
		workersCnt int
	}{
		{
			workersCnt: 1,
			tasksCnt:   10,
		},
		{
			workersCnt: 2,
			tasksCnt:   21,
		},
		{
			workersCnt: 10,
			tasksCnt:   502,
		},
	}
	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(tt *testing.T) {
			test := test
			pipe := pools.NewBusPool(test.workersCnt, 0, nil)
			for i := 0; i <= test.tasksCnt; i++ {
				task := &mocks.Task{}
				task.On("Do").After(time.Duration(rand.Intn(10)) * time.Millisecond).Once().Return(nil)
				defer task.AssertExpectations(tt)
				assert.NoError(tt, pipe.Push(task))
			}

			pipe.Wait()
			assert.NoError(tt, pipe.Close())
		})
	}
}

func Test_Pool_Errors(t *testing.T) {
	pipe := pools.NewBusPool(1, 0, nil)
	assert.NoError(t, pipe.Close())
	assert.EqualError(t, pipe.Close(), "pool close err: close of closed channel")
	assert.EqualError(t, pipe.Push(nil), "pool push err: send on closed channel")
}

func Test_Pool_NoThreads(t *testing.T) {
	assert.PanicsWithValue(t, "BusPool requires at least 1 thread", func() {
		pools.NewBusPool(0, 0, nil)
	})
}
