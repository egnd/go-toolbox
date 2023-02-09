package telebotmdw

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/telebot.v3"
)

const (
	ctxLoggerKey = "logger"
)

// CtxGetLogger return logger from telebot context.
func CtxGetLogger(ctx telebot.Context) (zerolog.Logger, bool) {
	res, ok := ctx.Get(ctxLoggerKey).(zerolog.Logger)

	return res, ok
}

// CtxWithLogger put logger into telebot context.
func CtxWithLogger(ctx telebot.Context, logger zerolog.Logger) telebot.Context {
	ctx.Set(ctxLoggerKey, logger)

	return ctx
}

// AddCtxLogger ad logger to telebot context.
func AddCtxLogger(logger *zerolog.Logger) telebot.MiddlewareFunc { //nolint:cyclop
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) error {
			ctxLog := logger.With()

			switch {
			case ctx.Update().Message != nil:
				ctxLog = ctxLog.Str("event", "message")
			case ctx.Update().EditedMessage != nil:
				ctxLog = ctxLog.Str("event", "edited_message")
			case ctx.Update().Callback != nil:
				ctxLog = ctxLog.Str("event", "callback")
				if ctx.Update().Callback.Unique != "" {
					ctxLog = ctxLog.Str("unique", ctx.Update().Callback.Unique)
				}
			case ctx.Update().Query != nil:
				ctxLog = ctxLog.Str("event", "query")
			case ctx.Update().InlineResult != nil:
				ctxLog = ctxLog.Str("event", "inline")
			}

			if ctx.Data() != "" {
				ctxLog = ctxLog.Str("data", ctx.Data())
			}

			if ctx.Text() != "" {
				ctxLog = ctxLog.Str("text", ctx.Text())
			}

			if args := strings.Join(ctx.Args(), ","); args != "" {
				ctxLog = ctxLog.Str("args", args)
			}

			sender := strings.TrimSpace(fmt.Sprintf("%s %s", ctx.Sender().FirstName, ctx.Sender().LastName))
			if sender == "" {
				ctxLog = ctxLog.Str("sender", ctx.Sender().Username)
			} else {
				ctxLog = ctxLog.Str("sender", fmt.Sprintf("[%s] %s", ctx.Sender().Username, sender))
			}

			return next(CtxWithLogger(ctx, ctxLog.Logger()))
		}
	}
}

// LogUpdate logs incoming update.
func LogUpdate(details bool) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) error {
			if details {
				data, _ := json.MarshalIndent(ctx.Update(), "", "  ")
				log.Println(string(data))
			}

			if err := next(ctx); err != nil {
				return err
			}

			if logger, ok := CtxGetLogger(ctx); ok {
				logger.Debug().Msg("update")
			}

			return nil
		}
	}
}
