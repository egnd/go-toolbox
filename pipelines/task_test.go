package pipelines_test

import (
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/pipelines"
	"github.com/egnd/go-toolbox/pipelines/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NewTaskExecutor(t *testing.T) {
	cases := []string{
		"hello world",
		"",
	}

	for k, phrase := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			var res string
			decorators := []pipelines.TaskDecorator{}
			if len(phrase) > 0 {
				for i := 0; i < len(phrase)-1; i++ {
					i := i
					decorators = append(decorators, func(next pipelines.TaskExecutor) pipelines.TaskExecutor {
						return func(task pipelines.Task) error {
							res += phrase[i : i+1]
							return next(task)
						}
					})
				}
			}

			task := &mocks.Task{}
			task.On("Do").Return(nil).Run(func(_ mock.Arguments) {
				if len(phrase) > 0 {
					res += phrase[len(phrase)-1:]
				}
			})

			pipelines.NewTaskExecutor(decorators)(task)

			assert.EqualValues(t, phrase, res+"")
		})
	}
}
