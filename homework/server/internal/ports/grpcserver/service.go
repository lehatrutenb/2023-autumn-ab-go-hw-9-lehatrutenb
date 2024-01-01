package grpcserver

import (
	"context"
	"homework/server/internal/ports/grpcserver/filemsg"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const MaxLoggedDataLen = 15

// мне не нравятся названия функций, но я пока не придумал как сделать лучше
// кажется, что было бы сильно проще использовать вместо oneof дополнительный message для GetFileInfo

func (s *server) GetFileData(stream filemsg.FileService_GetFileDataServer) error {
	in, err := stream.Recv()
	if err == io.EOF { // тут хочется вернуть ошибку, но я не нашёл прямо 1 в 1 по значению
		s.logger.Warn("End of stream without using it")
		return status.Errorf(codes.InvalidArgument, "End of stream without using it")
	}

	if err != nil {
		s.logger.Fatal("Failed to receive a message", zap.Error(err))
		return err
	}

	switch in.Param.(type) {
	case *filemsg.FileRequest_Name:
	default:
		s.logger.Error("Expected request with field Name")
		return status.Errorf(codes.InvalidArgument, "Expected request with field Name")
	}

	s.logger.Info("Get file data request",
		zap.String("fileName", in.GetName()))

	name := in.GetName()
	data, err := s.getFileDividedData(&s.app, name)
	if err != nil {
		return status.Errorf(codes.Internal, "Error occured in repo: %v", err)
	}

	err = s.sendData(stream, data)
	if err != nil {
		return err
	}

	if len(*data) <= 1 && len((*data)[0]) < MaxLoggedDataLen {
		s.logger.Info("Send data response",
			zap.Int("Size", len((*data)[0])), zap.Binary("Data", (*data)[0]))
	} else {
		// every chunk has simular size, that is why whole size is
		// (size of first chunk) * (amount of chunks without last) + size of last chunk
		s.logger.Info("Send data response",
			zap.Int("Size", len((*data)[0])*(len(*data)-1)+len((*data)[len(*data)-1])))
	}

	return nil
}

func (s *server) GetFileInfo(ctx context.Context, in *filemsg.FileRequest) (*filemsg.FileInfoResponse, error) {
	switch in.Param.(type) {
	case *filemsg.FileRequest_Name:
	default:
		s.logger.Error("Expected request with field Name")
		return nil, status.Errorf(codes.InvalidArgument, "Expected request with field Name")
	}

	s.logger.Info("Get file info request",
		zap.String("fileName", in.GetName()))

	fi, err := s.getFileInfo(&s.app, in.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occured in repo: %v", err)
	}

	s.logger.Info("Send file info response",
		zap.String("Name", fi.Name), zap.Int64("Size", fi.Size), zap.Uint32("Mode", fi.Mode), zap.Int64("ModTime", fi.Time))

	return fi, nil
}

func (s *server) GetFileNames(ctx context.Context, _ *emptypb.Empty) (*filemsg.FileListResponse, error) {
	s.logger.Info("Get file names request")

	names, err := s.getFileNames(&s.app)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occured in repo: %v", err)
	}

	s.logger.Info("Send file names response", zap.Strings("Names", names.Name))

	return names, nil
}
