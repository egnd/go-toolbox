// Package telebotmdw contains telebot middlwares.
package telebotmdw

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

// UsersWhiteList checks that sender is from whitelist.
func UsersWhiteList(usernames []string) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) error {
			for _, username := range usernames {
				if username == ctx.Sender().Username {
					if ctx.Callback() != nil {
						defer ctx.Respond() //nolint:errcheck
					}

					return next(ctx)
				}
			}

			return fmt.Errorf("UsersWhiteList error: user %s not found", ctx.Sender().Username)
		}
	}
}
