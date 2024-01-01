package grpcserver

import (
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver/filemsg"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	filemsg.UnimplementedFileServiceServer
	app    app.App
	logger *zap.Logger
}

func NewServer(a app.App, lr *zap.Logger) *grpc.Server {
	s := grpc.NewServer()
	filemsg.RegisterFileServiceServer(s, &server{app: a, logger: lr})

	return s
}
