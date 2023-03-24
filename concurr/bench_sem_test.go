package concurr_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/pipelines/pools"
	"github.com/egnd/go-toolbox/pipelines/tasks"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func Benchmark_Sem(b *testing.B) {
	for _, cnt := range []int{1, 10, 100, 1000} {
		ctx := context.Background()

		b.Run(fmt.Sprintf("sync/errgroup_%d", cnt), func(b *testing.B) {
			sem, _ := errgroup.WithContext(ctx)
			sem.SetLimit(cnt)
			for k := 0; k < b.N; k++ {
				sem.Go(func() error { return nil })
			}
			sem.Wait()
		})

		b.Run(fmt.Sprintf("sync/semaphore_%d", cnt), func(b *testing.B) {
			sem := semaphore.NewWeighted(int64(cnt))
			for k := 0; k < b.N; k++ {
				if err := sem.Acquire(ctx, 1); err != nil {
					b.Error(err)
					return
				}
				go func() error {
					defer sem.Release(1)
					return nil
				}()
			}
		})

		b.Run(fmt.Sprintf("go-toolbox_%d", cnt), func(b *testing.B) {
			sem := pools.NewSemaphore(cnt, nil)
			defer sem.Close()

			for k := 0; k < b.N; k++ {
				sem.Push(tasks.NewFunc("asdf", func() error { return nil }))
			}

			sem.Wait()
		})

		fmt.Println()
	}
}
