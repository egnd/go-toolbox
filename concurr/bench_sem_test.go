package concurr_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/concurr"
	sem "github.com/egnd/go-toolbox/concurr/semaphore"
	"github.com/egnd/go-toolbox/pipelines/pools"
	"github.com/egnd/go-toolbox/pipelines/tasks"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func Benchmark_Sem(b *testing.B) {
	for _, cnt := range []int{1, 10, 100, 1000} {
		// errgroup/semaphore
		bench_errgroup(cnt, b)

		// sync/semaphore
		bench_sync_sem(cnt, b)

		// pipelines/semaphore
		bench_pipe_sem(cnt, b)

		// concurr/semaphore
		bench_concurr_sem(cnt, b)

		fmt.Println()
	}
}

func bench_errgroup(limit int, b *testing.B) {
	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)
	group.SetLimit(limit)

	b.Run(fmt.Sprintf("sync/errgroup_%d", limit), func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			group.Go(func() error {
				time.Sleep(25 * time.Millisecond)

				return nil
			})
		}

		group.Wait()
	})
}

func bench_sync_sem(limit int, b *testing.B) {
	ctx := context.Background()
	syncsem := semaphore.NewWeighted(int64(limit))

	b.Run(fmt.Sprintf("sync/semaphore_%d", limit), func(b *testing.B) {
		var wg sync.WaitGroup
		for k := 0; k < b.N; k++ {
			if err := syncsem.Acquire(ctx, 1); err != nil {
				b.Error(err)
				return
			}

			wg.Add(1)

			go func() error {
				defer func() {
					syncsem.Release(1)
					wg.Done()
				}()

				time.Sleep(25 * time.Millisecond)

				return nil
			}()
		}

		wg.Wait()
	})
}
func bench_pipe_sem(limit int, b *testing.B) {
	pipesem := pools.NewSemaphore(limit, nil)
	defer pipesem.Close()

	b.Run(fmt.Sprintf("pipelines_sem_%d", limit), func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			pipesem.Push(tasks.NewFunc("asdf", func() error {
				time.Sleep(25 * time.Millisecond)

				return nil
			}))
		}

		pipesem.Wait()
	})
}

func bench_concurr_sem(limit int, b *testing.B) {
	ctx := context.Background()
	concsem := sem.New(ctx, limit, nil)
	defer concsem.Close()

	b.Run(fmt.Sprintf("concurr_%d", limit), func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			concsem.Push(concurr.TaskFunc(func(ctx context.Context) error {
				time.Sleep(25 * time.Millisecond)

				return nil
			}))
		}

		concsem.Wait()
	})
}
