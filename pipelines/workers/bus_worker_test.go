package workers_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/pipelines"
	"github.com/egnd/go-toolbox/pipelines/mocks"
	"github.com/egnd/go-toolbox/pipelines/workers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_BusWorker_Do(t *testing.T) {
	cases := []struct {
		buffSize int
		tasksCnt int
	}{
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
		t.Run(fmt.Sprint(k), func(tt *testing.T) {
			executor := &mocks.TaskExecutor{}
			executor.On("Execute", mock.Anything).Times(test.tasksCnt).
				After(time.Duration(rand.Intn(10)) * time.Millisecond).
				Return(nil)

			bus := make(chan pipelines.Doer)
			var wg sync.WaitGroup
			wg.Add(test.tasksCnt)
			workers.NewBusWorker(bus, &wg, executor.Execute)
			for i := 0; i < test.tasksCnt; i++ {
				assert.NoError(tt, (<-bus).Do(nil))
			}
			wg.Wait()

			assert.NoError(tt, (<-bus).Close())
			executor.AssertExpectations(t)
		})
	}
}

func Test_BusWorker_Do_Error(t *testing.T) {
	bus := make(chan pipelines.Doer)
	workers.NewBusWorker(bus, &sync.WaitGroup{}, (&mocks.TaskExecutor{}).Execute)
	worker := <-bus
	worker.Close()
	assert.EqualError(t, worker.Do(nil), "worker do err: send on closed channel")
}

func Test_BusWorker_Close_Error(t *testing.T) {
	bus := make(chan pipelines.Doer)
	close(bus)
	w := workers.NewBusWorker(bus, &sync.WaitGroup{}, (&mocks.TaskExecutor{}).Execute)
	time.Sleep(50 * time.Millisecond)
	assert.EqualError(t, w.Close(), "worker close err: close of closed channel")
}

func Test_BusWorker_Panic(t *testing.T) {
	assert.PanicsWithValue(t, "worker requires WaitGroup", func() {
		workers.NewBusWorker(nil, nil, nil)
	})
}
