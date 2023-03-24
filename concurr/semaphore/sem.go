package semaphore

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/egnd/go-toolbox/concurr"
)

type semaphore struct {
	ctx         context.Context
	wg          *sync.WaitGroup
	limiter     chan struct{}
	taskWrapper concurr.TaskDecorator
}

func New(ctx context.Context,
	threadsCnt int, wg *sync.WaitGroup, decorators ...concurr.TaskDecorator,
) *semaphore {
	res := semaphore{
		ctx:     ctx,
		wg:      wg,
		limiter: make(chan struct{}, threadsCnt),
	}

	if wg == nil {
		res.wg = &sync.WaitGroup{}
	}

	if len(decorators) > 0 {
		res.taskWrapper = concurr.NewTaskDecorator(decorators...)
	}

	return &res
}

func (sem *semaphore) Push(task concurr.Task) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch val := r.(type) {
			case error:
				err = val
			case string:
				err = errors.New(val)
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	select {
	case sem.limiter <- struct{}{}:
	case <-sem.ctx.Done():
		return
	}

	sem.wg.Add(1)

	go func() {
		defer func() {
			<-sem.limiter
			sem.wg.Done()
		}()

		if sem.taskWrapper == nil {
			task.Do(sem.ctx)
		} else {
			sem.taskWrapper(task).Do(sem.ctx)
		}
	}()

	return err
}

func (sem *semaphore) Wait() error {
	sem.wg.Wait()

	return nil
}

func (sem *semaphore) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch val := r.(type) {
			case error:
				err = val
			case string:
				err = errors.New(val)
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	close(sem.limiter)

	return err
}
