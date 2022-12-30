package tgchain

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// IListener is a tg updates listener interface.
type IListener interface {
	Add(EventType, ...IEventHandler) IListener
	Listen(context.Context, tgbotapi.UpdatesChannel)
}

// @TODO: listener with workers pool

// Listener is struct which is listening and handling events.
type Listener struct {
	handlers map[EventType][]IEventHandler
	warnUpd  WarnUpd
	errUpd   ErrUpd
}

// NewListener constructor for Listener struct.
func NewListener(warnUpd WarnUpd, errUpd ErrUpd) *Listener {
	return &Listener{
		warnUpd:  warnUpd,
		errUpd:   errUpd,
		handlers: make(map[EventType][]IEventHandler),
	}
}

// Add adds decorators for handling specific Telegram event.
func (b *Listener) Add(event EventType, handlers ...IEventHandler) IListener { //nolint:ireturn
	b.handlers[event] = append(b.handlers[event], handlers...)

	return b
}

func (b *Listener) buildChains() map[EventType]IEventHandler {
	chains := map[EventType]IEventHandler{}

	for event, handlers := range b.handlers {
		chains[event] = nil

		for i := len(handlers) - 1; i >= 0; i-- {
			if chains[event] == nil {
				chains[event] = handlers[i]
			} else {
				handlers[i].Decorate(chains[event])
				chains[event] = handlers[i]
			}
		}
	}

	return chains
}

// Listen starts listening incoming messages in channel.
func (b *Listener) Listen(ctx context.Context, updChan tgbotapi.UpdatesChannel) {
	handlers := b.buildChains()

	for update := range updChan {
		event := GetEventFrom(update)

		if item, ok := handlers[event]; ok {
			if err := item.Handle(
				context.WithValue(ctx, CtxEventKey, event), update,
			); err != nil && b.errUpd != nil {
				b.errUpd("update", update, err)
			}

			continue
		}

		if b.warnUpd != nil {
			b.warnUpd("unexpected event", update)
		}
	}
}
