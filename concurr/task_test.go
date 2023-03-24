package concurr_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/concurr"
	"github.com/stretchr/testify/assert"
)

func ctxStrSliceSet(ctx context.Context, val []string) context.Context {
	return context.WithValue(ctx, "ctxStrSlice", val)
}

func ctxStrSliceGet(ctx context.Context) []string {
	if res, ok := ctx.Value("ctxStrSlice").([]string); ok {
		return res
	}

	return nil
}

func Test_NewTask(t *testing.T) {
	t.Parallel()

	cases := []struct {
		err        error
		res        []string
		decorators []concurr.TaskDecorator
	}{
		{
			err: errors.New("asdfg"),
			res: []string{"mdw1", "mdw2"},
			decorators: []concurr.TaskDecorator{
				func(next concurr.Task) concurr.Task {
					return concurr.TaskFunc(func(ctx context.Context) error {
						return next.Do(ctxStrSliceSet(ctx, append(ctxStrSliceGet(ctx), "mdw1")))
					})
				},
				func(next concurr.Task) concurr.Task {
					return concurr.TaskFunc(func(ctx context.Context) error {
						return next.Do(ctxStrSliceSet(ctx, append(ctxStrSliceGet(ctx), "mdw2")))
					})
				},
			},
		},
		{
			res: []string{"mdw10", "mdw20"},
			decorators: []concurr.TaskDecorator{
				func(next concurr.Task) concurr.Task {
					return concurr.TaskFunc(func(ctx context.Context) error {
						return next.Do(ctxStrSliceSet(ctx, append(ctxStrSliceGet(ctx), "mdw10")))
					})
				},
				func(next concurr.Task) concurr.Task {
					return concurr.TaskFunc(func(ctx context.Context) error {
						return next.Do(ctxStrSliceSet(ctx, append(ctxStrSliceGet(ctx), "mdw20")))
					})
				},
			},
		},
		{},
	}

	for k, test := range cases {
		k, test := k, test
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			t.Parallel()
			assert.EqualValues(t, test.err, concurr.NewTask(
				concurr.TaskFunc(func(ctx context.Context) error {
					assert.EqualValues(t, test.res, ctxStrSliceGet(ctx))
					return test.err
				}), test.decorators...,
			).Do(context.TODO()))
		})
	}
}

func Test_NewTaskDecorator(t *testing.T) {
	t.Parallel()

	cases := []struct {
		err        error
		res        []string
		decorators []concurr.TaskDecorator
	}{
		{
			err: errors.New("asdfg"),
			res: []string{"mdw1", "mdw2"},
			decorators: []concurr.TaskDecorator{
				func(next concurr.Task) concurr.Task {
					return concurr.TaskFunc(func(ctx context.Context) error {
						return next.Do(ctxStrSliceSet(ctx, append(ctxStrSliceGet(ctx), "mdw1")))
					})
				},
				func(next concurr.Task) concurr.Task {
					return concurr.TaskFunc(func(ctx context.Context) error {
						return next.Do(ctxStrSliceSet(ctx, append(ctxStrSliceGet(ctx), "mdw2")))
					})
				},
			},
		},
		{
			res: []string{"mdw1"},
			decorators: []concurr.TaskDecorator{
				func(next concurr.Task) concurr.Task {
					return concurr.TaskFunc(func(ctx context.Context) error {
						return next.Do(ctxStrSliceSet(ctx, append(ctxStrSliceGet(ctx), "mdw1")))
					})
				},
			},
		},
		{},
	}

	for k, test := range cases {
		k, test := k, test
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			t.Parallel()

			var decorate concurr.TaskDecorator

			if len(test.decorators) == 0 {
				assert.PanicsWithValue(t, "empty decorators list", func() { concurr.NewTaskDecorator(test.decorators...) })
				return
			} else {
				decorate = concurr.NewTaskDecorator(test.decorators...)
			}

			assert.EqualValues(t, test.err, decorate(concurr.TaskFunc(
				func(ctx context.Context) error {
					assert.EqualValues(t, test.res, ctxStrSliceGet(ctx))
					return test.err
				},
			)).Do(context.TODO()))
		})
	}
}
