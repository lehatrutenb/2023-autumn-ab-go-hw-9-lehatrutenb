package grpcserver

import (
	"context"
	"homework/server/internal/ports/grpcserver/filemsg"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// мне не нравятся названия функций, но я пока не придумал как сделать лучше

func (s *server) GetFileData(stream filemsg.FileService_GetFileDataServer) error {
	in, err := stream.Recv()
	if err == io.EOF { // тут хочется вернуть ошибку, но я не нашёл прямо 1 в 1 по значению
		return status.Errorf(codes.InvalidArgument, "End of stream without using it")
	}

	if err != nil {
		log.Fatalf("Failed to receive a message : %v", err)
		return err
	}

	switch in.Param.(type) {
	case *filemsg.FileRequest_Name:
	default:
		return status.Errorf(codes.InvalidArgument, "Expected request with Name")
	}

	name := in.GetName()
	data, err := GetFileDividedData(&s.App, name)
	if err != nil {
		return status.Errorf(codes.Internal, "Error occured in repo: %v", err)
	}

	err = sendData(stream, data)
	if err != nil {
		return err
	}

	return nil
}

func sendData(stream filemsg.FileService_GetFileDataServer, data *[][]byte) error {
	for _, ch := range *data { // go through every data chunk
		toSend := &filemsg.FileDataResponse{Data: ch}
		if err := stream.Send(toSend); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
			return err
		}

		err := checkIfRecieved(stream)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkIfRecieved(stream filemsg.FileService_GetFileDataServer) error {
	in, err := stream.Recv()
	if err == io.EOF {
		return status.Errorf(codes.InvalidArgument, "End of stream without using it")
	}
	if err != nil {
		return err
	}

	switch in.Param.(type) {
	case *filemsg.FileRequest_Got:
	default:
		return status.Errorf(codes.InvalidArgument, "Expected request with Got")
	}
	return nil
}

func (s *server) GetFileInfo(ctx context.Context, in *filemsg.FileRequest) (*filemsg.FileInfoResponse, error) {
	switch in.Param.(type) {
	case *filemsg.FileRequest_Name:
	default:
		return nil, status.Errorf(codes.InvalidArgument, "Expected request with Name")
	}

	fInf, err := GetFileInfo(&s.App, in.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occured in repo: %v", err)
	}

	return fInf, nil
}

func (s *server) GetFileNames(ctx context.Context, _ *emptypb.Empty) (*filemsg.FileListResponse, error) {
	names, err := GetFileNames(&s.App)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occured in repo: %v", err)
	}

	return names, nil
}
