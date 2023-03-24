package concurr_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/concurr"
)

func Benchmark_Tasks(b *testing.B) {
	ctxSet := func(ctx context.Context, val int) context.Context {
		return context.WithValue(ctx, "ctxSet", val)
	}

	ctxGet := func(ctx context.Context) int {
		if res, ok := ctx.Value("ctxSet").(int); ok {
			return res
		}

		return 0
	}

	benchTask := concurr.TaskFunc(func(ctx context.Context) error {
		ctxGet(ctx)
		return nil
	})

	benchDecorator := func(next concurr.Task) concurr.Task {
		return concurr.TaskFunc(func(ctx context.Context) error {
			return next.Do(ctxSet(ctx, ctxGet(ctx)))
		})
	}

	for _, cnt := range []int{1, 10, 100, 1000} {
		ctx := ctxSet(context.Background(), cnt)
		decorators := make([]concurr.TaskDecorator, cnt)
		for num := range decorators {
			decorators[num] = benchDecorator
		}

		// decorate task with builded chain
		b.Run(fmt.Sprintf("NewTaskDecorator_%d", cnt), func(b *testing.B) {
			chain := concurr.NewTaskDecorator(decorators...)
			for k := 0; k < b.N; k++ {
				chain(benchTask).Do(ctx)
			}
		})

		// decorate task during execution
		b.Run(fmt.Sprintf("NewTask_%d", cnt), func(b *testing.B) {
			for k := 0; k < b.N; k++ {
				concurr.NewTask(benchTask, decorators...).Do(ctx)
			}
		})

		fmt.Println()
	}
}
