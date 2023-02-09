package tgchain

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// IEventHandler is event handler interface.
type IEventHandler interface {
	GetName() string
	GetDescription() string
	Decorate(IEventHandler)
	Next(context.Context, tgbotapi.Update) error
	Handle(context.Context, tgbotapi.Update) error
}

// AbstractHandler is a abstract event handler.
type AbstractHandler struct {
	next IEventHandler
}

// GetName returns handler name.
func (h *AbstractHandler) GetName() string { return "" }

// GetDescription returns handler description.
func (h *AbstractHandler) GetDescription() (descr string) { return "" }

// Decorate is a setter for next handler.
func (h *AbstractHandler) Decorate(next IEventHandler) {
	h.next = next
}

// Next pass update to next handler in chain.
func (h *AbstractHandler) Next(ctx context.Context, upd tgbotapi.Update) error {
	if h.next == nil {
		return nil
	}

	return h.next.Handle(ctx, upd) //nolint:wrapcheck
}

// ReplyToMsg sends reply to tg message.
func (h *AbstractHandler) ReplyToMsg(msg *tgbotapi.Message, text string, api ITgAPI) (err error) {
	if msg == nil || len(text) == 0 {
		return
	}

	resp := tgbotapi.NewMessage(msg.Chat.ID, text)
	resp.ReplyToMessageID = msg.MessageID

	_, err = api.Send(resp)

	return
}
