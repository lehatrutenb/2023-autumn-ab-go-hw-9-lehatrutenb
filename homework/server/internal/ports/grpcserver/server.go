package grpcserver

import (
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver/filemsg"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	filemsg.UnimplementedFileServiceServer
	app              app.App
	logger           *zap.Logger
	MaxLoggedDataLen int
}

type ServerOption func(*server)

func NewServer(a app.App, lr *zap.Logger, addr string, opts ...ServerOption) (net.Listener, *grpc.Server) {
	grpcS := grpc.NewServer()
	server := &server{app: a, logger: lr, MaxLoggedDataLen: 15}

	for _, opt := range opts {
		opt(server)
	}

	filemsg.RegisterFileServiceServer(grpcS, server)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis, grpcS
}

func WithMaxLoggedDataLen(maxDataLen int) ServerOption {
	return func(s *server) {
		s.MaxLoggedDataLen = maxDataLen
	}
}
