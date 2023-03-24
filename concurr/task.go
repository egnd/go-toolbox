package concurr

import (
	"context"
)

type TaskFunc func(context.Context) error

func (task TaskFunc) Do(ctx context.Context) error {
	return task(ctx)
}

type TaskDecorator func(next Task) Task

func NewTaskDecorator(decorators ...TaskDecorator) TaskDecorator {
	switch len(decorators) {
	case 0:
		panic("empty decorators list")
	case 1:
		return decorators[0]
	}

	res := decorators[len(decorators)-1]

	for i := len(decorators) - 2; i >= 0; i-- {
		i, old := i, res
		res = func(next Task) Task {
			return decorators[i](old(next))
		}
	}

	return res
}

type Task interface {
	Do(ctx context.Context) error
}

func NewTask(task Task, decorators ...TaskDecorator) Task {
	for i := len(decorators) - 1; i >= 0; i-- {
		task = decorators[i](task)
	}

	return task
}
