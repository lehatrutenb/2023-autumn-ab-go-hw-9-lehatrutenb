package grpcclient

import (
	"go.uber.org/zap"
)

// не хочу выносить адресс сервера в структуру, тк почему бы ему не меняться
type client struct {
	logger *zap.Logger
}

func NewClient(lr *zap.Logger) *client {
	return &client{logger: lr}
}
