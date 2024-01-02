package grpcclient

import (
	"time"

	"go.uber.org/zap"
)

// не хочу выносить адресс сервера в структуру, тк почему бы ему не меняться
type Client struct {
	logger           *zap.Logger
	RequestTimeout   time.Duration
	StreamTimeout    time.Duration
	MaxLoggedDataLen int
}

type ClientOption func(*Client)

func NewClient(lr *zap.Logger, opts ...ClientOption) *Client {
	c := &Client{logger: lr, RequestTimeout: time.Second, StreamTimeout: 10 * time.Second, MaxLoggedDataLen: 15}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithRequestTimeout(rt time.Duration) ClientOption {
	return func(c *Client) {
		c.RequestTimeout = rt
	}
}

func WithStreamTimeout(st time.Duration) ClientOption {
	return func(c *Client) {
		c.StreamTimeout = st
	}
}

func WithMaxLoggedDataLen(maxDataLen int) ClientOption {
	return func(c *Client) {
		c.MaxLoggedDataLen = maxDataLen
	}
}
