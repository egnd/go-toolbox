package decorators_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/pipelines"
	"github.com/egnd/go-toolbox/pipelines/decorators"
	"github.com/egnd/go-toolbox/pipelines/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_CatchPanic(t *testing.T) {
	cases := []struct {
		panic string
	}{
		{"123"},
		{"dv78v9s"},
		{},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			task := &mocks.Task{}
			if test.panic == "" {
				task.On("Do").Return(nil)
				pipelines.NewTaskExecutor([]pipelines.TaskDecorator{func(next pipelines.TaskExecutor) pipelines.TaskExecutor {
					return func(task pipelines.Task) (err error) {
						assert.NoError(t, next(task))
						return
					}
				}, decorators.CatchPanic})(task)
			} else {
				task.On("Do").Panic(test.panic)
				pipelines.NewTaskExecutor([]pipelines.TaskDecorator{func(next pipelines.TaskExecutor) pipelines.TaskExecutor {
					return func(task pipelines.Task) (err error) {
						assert.EqualValues(t, errors.New(test.panic), next(task))
						return
					}
				}, decorators.CatchPanic})(task)
			}
		})
	}
}

func Test_ThrowPanic(t *testing.T) {
	cases := []struct {
		err error
	}{
		{},
		{errors.New("test error")},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			task := &mocks.Task{}
			task.On("Do").Return(test.err)

			if test.err == nil {
				pipelines.NewTaskExecutor([]pipelines.TaskDecorator{decorators.ThrowPanic})(task)
			} else {
				assert.PanicsWithError(t, test.err.Error(), func() {
					pipelines.NewTaskExecutor([]pipelines.TaskDecorator{decorators.ThrowPanic})(task)
				})
			}
		})
	}
}
