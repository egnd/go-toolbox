package tasks_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/pipelines/tasks"

	"github.com/stretchr/testify/assert"
)

func Test_FuncTask(t *testing.T) {
	cases := []struct {
		id   string
		task func() error
		err  error
	}{
		{
			id: "asdfg",
			task: func() error {
				return nil
			},
		},
		{
			task: func() error {
				return errors.New("error")
			},
			err: errors.New("error"),
		},
		{},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			task := tasks.NewFunc(test.id, test.task)
			assert.EqualValues(t, test.id, task.ID())
			assert.EqualValues(t, test.err, task.Do())
		})
	}
}
