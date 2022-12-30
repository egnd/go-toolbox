// Package decorators contains decorators for pipeline tasks
package decorators

import (
	"github.com/egnd/go-toolbox/pipelines"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

// LogErrorZero logs task error with zerolog logger.
func LogErrorZero(logger zerolog.Logger) pipelines.TaskDecorator {
	return func(next pipelines.TaskExecutor) pipelines.TaskExecutor {
		return func(task pipelines.Task) (err error) {
			if err = next(task); err != nil {
				logger.Error().Err(err).Str("task", task.ID()).Msg("do")
			}

			return
		}
	}
}

// LogErrorZap logs task error with zap logger.
func LogErrorZap(logger *zap.Logger) pipelines.TaskDecorator {
	return func(next pipelines.TaskExecutor) pipelines.TaskExecutor {
		return func(task pipelines.Task) (err error) {
			if err = next(task); err != nil {
				logger.Error("do", zap.Error(err), zap.String("task", task.ID()))
			}

			return
		}
	}
}
