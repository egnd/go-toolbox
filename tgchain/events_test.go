package tgchain_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/tgchain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
)

func Test_GetEventFromCtx(t *testing.T) {
	assert.EqualValues(t, tgchain.EventMessage,
		tgchain.GetEventFromCtx(
			context.WithValue(context.Background(), tgchain.CtxEventKey, tgchain.EventMessage),
		),
	)
}

func Test_GetEventFrom(t *testing.T) {
	cases := []struct {
		upd   tgbotapi.Update
		event tgchain.EventType
	}{
		{
			event: tgchain.EventUndefined,
		},
		{
			upd:   tgbotapi.Update{InlineQuery: &tgbotapi.InlineQuery{}},
			event: tgchain.EventInlineQuery,
		},
		{
			upd:   tgbotapi.Update{CallbackQuery: &tgbotapi.CallbackQuery{}},
			event: tgchain.EventCallbackQuery,
		},
		{
			upd:   tgbotapi.Update{Message: &tgbotapi.Message{Entities: &[]tgbotapi.MessageEntity{{Type: "bot_command"}}}},
			event: tgchain.EventCommand,
		},
		{
			upd:   tgbotapi.Update{Message: &tgbotapi.Message{}},
			event: tgchain.EventMessage,
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k), func(tt *testing.T) {
			assert.EqualValues(tt, test.event, tgchain.GetEventFrom(test.upd))
		})
	}
}

func Test_GetEventName(t *testing.T) {
	cases := []struct {
		event tgchain.EventType
		name  string
	}{
		{event: tgchain.EventMessage, name: "message"},
		{event: tgchain.EventCommand, name: "command"},
		{event: tgchain.EventInlineQuery, name: "inline_query"},
		{event: tgchain.EventCallbackQuery, name: "callback_query"},
		{event: tgchain.EventUndefined, name: "undefined"},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k), func(tt *testing.T) {
			assert.EqualValues(tt, test.name, tgchain.GetEventName(test.event))
		})
	}
}
