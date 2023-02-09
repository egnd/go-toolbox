package telebotmdw

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

// CatchPanics catch panic middlware.
func CatchPanics(catch bool) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) (err error) {
			if !catch {
				return next(ctx)
			}

			defer func() {
				if panicInfo := recover(); panicInfo != nil {
					switch val := panicInfo.(type) {
					case error:
						err = fmt.Errorf("panic: %w", val)
					case string:
						err = fmt.Errorf("panic: %s", val)
					default:
						err = fmt.Errorf("panic: %v", val)
					}
				}
			}()

			err = next(ctx)

			return
		}
	}
}
