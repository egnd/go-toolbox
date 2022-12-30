package decorators

import (
	"fmt"

	"github.com/egnd/go-toolbox/pipelines"
)

// CatchPanic is catching panic and return it as an error.
func CatchPanic(next pipelines.TaskExecutor) pipelines.TaskExecutor {
	return func(task pipelines.Task) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
			}
		}()

		err = next(task)

		return
	}
}

// ThrowPanic is throws task error as a panic.
func ThrowPanic(next pipelines.TaskExecutor) pipelines.TaskExecutor {
	return func(task pipelines.Task) error {
		if err := next(task); err != nil {
			panic(err)
		}

		return nil
	}
}
