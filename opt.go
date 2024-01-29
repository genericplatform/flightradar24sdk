package flightradar24sdk

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type apiOptions struct {
	client       *http.Client
	logger       resty.Logger
	debugEnabled bool
}

type Option func(o *apiOptions)

func WithClient(client *http.Client) Option {
	return func(o *apiOptions) {
		o.client = client
	}
}

func WithLogger(logger resty.Logger) Option {
	return func(o *apiOptions) {
		o.logger = logger
	}
}

func WithDebug(enabled bool) Option {
	return func(o *apiOptions) {
		o.debugEnabled = enabled
	}
}
