// Package workers contains different types of workers
package workers

import (
	"errors"
	"fmt"
	"sync"

	"github.com/egnd/go-toolbox/pipelines"
)

// BusWorker struct for handling tasks.
type BusWorker struct {
	tasks chan pipelines.Task
	stop  chan struct{}
}

// NewBusWorker creates workers for pool pipelines.
func NewBusWorker(
	bus chan<- pipelines.Doer, wg *sync.WaitGroup, execute pipelines.TaskExecutor, //nolint:varnamelen
) *BusWorker {
	if wg == nil {
		panic("worker requires WaitGroup")
	}

	worker := &BusWorker{
		tasks: make(chan pipelines.Task, 1),
		stop:  make(chan struct{}),
	}

	go func() {
		for {
			if err := worker.notify(bus); err != nil {
				return
			}

			for task := range worker.tasks {
				_ = execute(task)

				wg.Done()

				break
			}
		}
	}()

	return worker
}

func (w *BusWorker) notify(bus chan<- pipelines.Doer) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = w.Close()
		}
	}()

	select {
	case bus <- w:
	case <-w.stop:
		err = errors.New("worker is stopped")
	}

	return
}

// Do is putting task to worker's queue.
func (w *BusWorker) Do(task pipelines.Task) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("worker do err: %v", r)
		}
	}()

	w.tasks <- task

	return
}

// Close is stopping a worker.
func (w *BusWorker) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("worker close err: %v", r)
		}
	}()

	close(w.tasks)
	w.stop <- struct{}{}
	close(w.stop)

	return
}
