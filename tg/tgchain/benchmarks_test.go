package tgchain_test

import (
	"context"
	"sync"
	"testing"

	"github.com/egnd/go-toolbox/tg/tgchain"
	"github.com/egnd/go-toolbox/tg/tgchain/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/mock"
)

var (
	ctx = context.Background()
)

func Benchmark_Listener(b *testing.B) {
	updChan := make(chan tgbotapi.Update)

	handler := &mocks.IEventHandler{}
	handler.On("Handle", mock.Anything, mock.Anything).Return(nil)

	l := tgchain.NewListener(nil, nil)
	l.Add(tgchain.EventMessage, handler)

	var wd sync.WaitGroup
	wd.Add(1)
	go func() {
		defer wd.Done()
		l.Listen(ctx, updChan)
	}()

	for i := 0; i < b.N; i++ {
		updChan <- tgbotapi.Update{Message: &tgbotapi.Message{}}
	}

	close(updChan)
	wd.Wait()
}
