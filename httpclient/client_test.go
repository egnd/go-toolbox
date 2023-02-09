package httpclient_test

import (
	"testing"
	"time"

	"github.com/egnd/go-toolbox/httpclient"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_NewHTTPClient(t *testing.T) {
	cfg := viper.New()
	cfg.Set("timeout", "10s")
	client := httpclient.NewHTTPClient(cfg, nil)

	assert.EqualValues(t, 10*time.Second, client.Timeout)
	assert.Empty(t, client.Transport)
}
