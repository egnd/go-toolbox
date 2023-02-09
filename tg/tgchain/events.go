package tgchain

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// EventType variable type for event type ID.
type EventType int

const (
	// EventUndefined undefined event.
	EventUndefined EventType = iota
	// EventMessage receiving message event.
	EventMessage
	// EventCommand receiving command event.
	EventCommand
	// EventInlineQuery receiving inline query event.
	EventInlineQuery
	// EventCallbackQuery receiving callback query event.
	EventCallbackQuery
)

// CtxEventKeyType is a type for event name in context.
type CtxEventKeyType int

const (
	// CtxEventKey context key for storing event type ID.
	CtxEventKey CtxEventKeyType = iota
)

// GetEventFromCtx return event ID from context struct.
func GetEventFromCtx(ctx context.Context) EventType {
	return ctx.Value(CtxEventKey).(EventType) //nolint:forcetypeassert
}

// GetEventFrom return event type from tg update.
func GetEventFrom(upd tgbotapi.Update) EventType {
	switch {
	case upd.Message != nil && !upd.Message.IsCommand():
		return EventMessage
	case upd.Message != nil && upd.Message.IsCommand():
		return EventCommand
	case upd.InlineQuery != nil:
		return EventInlineQuery
	case upd.CallbackQuery != nil:
		return EventCallbackQuery
	default:
		return EventUndefined
	}
}

// GetEventName return event name by it's ID.
func GetEventName(eventID EventType) string {
	switch eventID {
	case EventMessage:
		return "message"
	case EventCommand:
		return "command"
	case EventInlineQuery:
		return "inline_query"
	case EventCallbackQuery:
		return "callback_query"
	case EventUndefined:
		fallthrough
	default:
		return "undefined"
	}
}
