package tgchain_test

import (
	"context"
	"errors"
	"testing"

	"github.com/egnd/go-toolbox/tg/tgchain"
	"github.com/egnd/go-toolbox/tg/tgchain/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
)

func Test_AbstractHandler(t *testing.T) {
	ctx := context.Background()
	upd := tgbotapi.Update{}

	mHandler := &mocks.IEventHandler{}
	mHandler.On("Handle", ctx, upd).Return(nil)

	h := tgchain.AbstractHandler{}
	assert.EqualValues(t, "", h.GetName())
	assert.EqualValues(t, "", h.GetDescription())
	h.Next(ctx, upd)

	h.Decorate(mHandler)
	h.Next(ctx, upd)

	mHandler.AssertExpectations(t)

	api := &mocks.ITgAPI{}

	assert.NoError(t, h.ReplyToMsg(nil, "", api))

	msg := &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: 1}}
	resp := tgbotapi.NewMessage(msg.Chat.ID, "text")
	resp.ReplyToMessageID = msg.MessageID
	api.On("Send", resp).Return(tgbotapi.Message{}, nil).Once()
	api.On("Send", resp).Return(tgbotapi.Message{}, errors.New("error")).Once()

	assert.NoError(t, h.ReplyToMsg(msg, "text", api))
	assert.Error(t, h.ReplyToMsg(msg, "text", api))
}
