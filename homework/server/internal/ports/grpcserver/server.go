package grpcserver

import (
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver/filemsg"

	"google.golang.org/grpc"
)

type server struct {
	filemsg.UnimplementedFileServiceServer
	App app.App
}

func NewServer(a app.App) *grpc.Server {
	s := grpc.NewServer()
	filemsg.RegisterFileServiceServer(s, &server{App: a})

	return s
}
