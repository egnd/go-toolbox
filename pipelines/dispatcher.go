// Package pipelines contains tools for parallel execution
package pipelines

import (
	"io"
)

// Dispatcher is a pool interface.
type Dispatcher interface {
	io.Closer
	Push(Task) error
	Wait()
}
