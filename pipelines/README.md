# pipelines

Golang package for parallel execution of tasks.

### Pools types:
* BusPool: Common Pool of Workers. The Task is taken into work by the first released Worker.
* HashPool: Worker pool, which allows you to change the strategy for assigning Tasks to Workers.
* Semaphore: Primitive for limiting the number of threads for the Tasks parallel execution.

### Examples:
```golang
package main

import (
	"fmt"
	"sync"

	"github.com/egnd/go-toolbox/pipelines/assign"
	"github.com/egnd/go-toolbox/pipelines/decorators"
	"github.com/egnd/go-toolbox/pipelines/pools"
	"github.com/egnd/go-toolbox/pipelines/tasks"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

func main() {
	var wg sync.WaitGroup

	// BusPool example:
	pipe := pools.NewBusPool(
		2,  // set parallel threads count
		10, // set tasks queue size
		&wg,
		// add some task decorators:
		decorators.LogErrorZero(zerolog.Nop()), // log tasks errors with zerolog logger
		decorators.CatchPanic,                  // convert tasks panics to errors
	)

	// HashPool example:
	pipe := pools.NewHashPool(
		2,  // set parallel threads count
		10, // set tasks queue size
		&wg,
		assign.Sticky, // choose tasks to workers assignment method
		// add some task decorators:
		decorators.LogErrorZap(zap.NewNop()), // log tasks errors with zap logger
		decorators.CatchPanic,                // convert tasks panics to errors
	)

	// Semaphore example:
	pipe := pools.NewSemaphore(2, // set parallel threads count
		&wg,
		// add some task decorators:
		decorators.ThrowPanic, // convert tasks errors to panics
	)

	// Send some tasks to pool
	for i := 0; i < 10; i++ {
		pipe.Push(tasks.NewFunc("task#"+fmt.Sprint(i), func() (err error) {
			// do something
			return err
		}))
	}

	// Wait for task processing
	pipe.Wait()

	// Close pool
	if err := pipe.Close(); err != nil {
		panic(err)
	}
}
```
