package sitemanager

import (
	"net/http"

	"github.com/murasame29/unifi-client-go/internal"
)

type Client struct {
	client *internal.Client
}

type Option = internal.Option

func WithBaseURL(baseURL string) Option {
	return internal.WithBaseURL(baseURL)
}

func WithHTTPClient(httpClient *http.Client) Option {
	return internal.WithHTTPClient(httpClient)
}

func NewClient(apiKey string, opts ...Option) *Client {
	return &Client{
		client: internal.NewClient(apiKey, opts...),
	}
}
