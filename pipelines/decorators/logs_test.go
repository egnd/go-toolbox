package decorators_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/pipelines"
	"github.com/egnd/go-toolbox/pipelines/decorators"
	"github.com/egnd/go-toolbox/pipelines/mocks"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

func Test_LogErrorZero(t *testing.T) {
	cases := []struct {
		taskID string
		err    error
	}{
		{
			taskID: "asdf",
			err:    errors.New("error"),
		},
		{
			taskID: "asdffff",
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			task := &mocks.Task{}
			task.On("ID").Return(test.taskID)
			task.On("Do").Return(test.err)

			decorators.LogErrorZero(zerolog.Nop())(
				func(task pipelines.Task) error { return task.Do() },
			)(task)
		})
	}
}

func Test_LogErrorZap(t *testing.T) {
	cases := []struct {
		taskID string
		err    error
	}{
		{
			taskID: "asdf",
			err:    errors.New("error"),
		},
		{
			taskID: "asdffff",
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			task := &mocks.Task{}
			task.On("ID").Return(test.taskID)
			task.On("Do").Return(test.err)

			decorators.LogErrorZap(zap.NewNop())(
				func(task pipelines.Task) error { return task.Do() },
			)(task)
		})
	}
}
