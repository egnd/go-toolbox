package workers_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/pipelines/mocks"
	"github.com/egnd/go-toolbox/pipelines/workers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Worker(t *testing.T) {
	cases := []struct {
		buffSize int
		tasksCnt int
	}{
		{
			tasksCnt: 15,
		},
		{
			buffSize: 10,
			tasksCnt: 21,
		},
		{
			buffSize: 20,
			tasksCnt: 10,
		},
	}
	for k, test := range cases {
		test := test
		t.Run(fmt.Sprint(k), func(tt *testing.T) {
			executor := &mocks.TaskExecutor{}
			executor.On("Execute", mock.Anything).Times(test.tasksCnt).
				After(time.Duration(rand.Intn(10)) * time.Millisecond).
				Return(nil)

			var wg sync.WaitGroup
			wg.Add(test.tasksCnt)
			worker := workers.NewWorker(test.buffSize, &wg, executor.Execute)
			for i := 0; i < test.tasksCnt; i++ {
				assert.NoError(tt, worker.Do(nil))
			}
			wg.Wait()

			assert.NoError(tt, worker.Close())
			executor.AssertExpectations(t)
		})
	}
}

func Test_Worker_Do_Error(t *testing.T) {
	worker := workers.NewWorker(0, &sync.WaitGroup{}, (&mocks.TaskExecutor{}).Execute)
	assert.NoError(t, worker.Close())
	assert.EqualError(t, worker.Do(nil), "worker do err: send on closed channel")
}

func Test_Worker_Close_Error(t *testing.T) {
	worker := workers.NewWorker(0, &sync.WaitGroup{}, (&mocks.TaskExecutor{}).Execute)
	assert.NoError(t, worker.Close())
	assert.EqualError(t, worker.Close(), "worker close err: close of closed channel")
}

func Test_Worker_Panic(t *testing.T) {
	assert.PanicsWithValue(t, "worker requires WaitGroup", func() {
		workers.NewWorker(0, nil, nil)
	})
}
