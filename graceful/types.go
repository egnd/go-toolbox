// Package graceful has tools for making graceful shutdown.
package graceful

import "errors"

var (
	ErrCallbackCnt    = errors.New("invalid callbacks count, required only start and stop callbacks") //nolint:revive
	ErrStartFuncType  = errors.New("invalid start func type")                                         //nolint:revive
	ErrEmptyStartFunc = errors.New("empty start func")                                                //nolint:revive
	ErrStopFuncType   = errors.New("invalid stop func type")                                          //nolint:revive
	ErrEmptyStopFunc  = errors.New("empty stop func")                                                 //nolint:revive
	ErrEmptyListener  = errors.New("undefined global listener")                                       //nolint:revive
)

// Callback is a common function type.
type Callback = func()

// CallbackWithError is a common function type with error.
type CallbackWithError = func() error
