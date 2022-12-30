// Package graceful has tools for making graceful shutdown.
package graceful

import "errors"

var (
	ErrCallbackCnt = errors.New("invalid callbacks count, required only start and stop callbacks") //nolint:revive
	ErrStartFunc   = errors.New("invalid start func")                                              //nolint:revive
	ErrStopFunc    = errors.New("invalid stop func")                                               //nolint:revive
)

// Callback is a common function type.
type Callback func() error
