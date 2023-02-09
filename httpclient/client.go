// Package httpclient contains tools for http clients.
package httpclient

import (
	"net/http"

	"github.com/spf13/viper"
)

// NewHTTPClient return http.Client.
func NewHTTPClient(cfg *viper.Viper, transport http.RoundTripper) *http.Client {
	return &http.Client{ //nolint:exhaustruct
		Timeout:   cfg.GetDuration("timeout"),
		Transport: transport,
		// Jar http.CookieJar @TODO:
	}
}
