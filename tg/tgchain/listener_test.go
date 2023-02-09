package tgchain_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/egnd/go-toolbox/tg/tgchain"
	"github.com/egnd/go-toolbox/tg/tgchain/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/mock"
)

func Test_Listener(t *testing.T) {
	cases := []struct {
		event    tgchain.EventType
		handlers []tgchain.IEventHandler
		upd      tgbotapi.Update
	}{
		{
			event: tgchain.EventMessage,
			upd:   tgbotapi.Update{Message: &tgbotapi.Message{}},
			handlers: func() []tgchain.IEventHandler {
				handler1 := &mocks.IEventHandler{}
				handler2 := &mocks.IEventHandler{}
				handler1.On("Decorate", handler2)
				handler1.On("Handle", mock.Anything, mock.Anything).Return(errors.New("error"))
				return []tgchain.IEventHandler{handler1, handler2}
			}(),
		},
		{
			event: tgchain.EventMessage,
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k), func(tt *testing.T) {
			updChan := make(chan tgbotapi.Update)
			warnUpd := &mocks.WarnUpd{}
			warnUpd.On("Execute", "unexpected event", test.upd).Maybe()
			errUpd := &mocks.ErrUpd{}
			errUpd.On("Execute", "update", test.upd, mock.Anything).Maybe()
			l := tgchain.NewListener(warnUpd.Execute, errUpd.Execute)
			l.Add(test.event, test.handlers...)
			var wd sync.WaitGroup
			wd.Add(1)
			go func() {
				defer wd.Done()
				l.Listen(context.TODO(), updChan)
			}()
			updChan <- test.upd
			close(updChan)
			wd.Wait()
			for _, mock := range test.handlers {
				mock.(*mocks.IEventHandler).AssertExpectations(tt)
			}
			warnUpd.AssertExpectations(tt)
		})
	}
}
