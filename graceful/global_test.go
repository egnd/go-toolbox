package graceful_test

import (
	"context"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/graceful"
	"github.com/stretchr/testify/assert"
)

func Test_Global(t *testing.T) {
	ctx, cancelFunc := graceful.Init(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancelFunc()
	}()

	assert.NoError(t, graceful.Register(func() { <-ctx.Done() }))
	assert.NoError(t, graceful.Wait())
}
