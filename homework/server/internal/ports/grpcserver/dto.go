package grpcserver

import (
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver/filemsg"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// dto funcs promise to log errors themselve

// TODO add optinos to change it
const SendByteSize = 1000

// Я не думаю, что прeобразвания в типы сообщений из protoc надо делать в app, так что думаю тут это хорошо

func (s *server) getFileInfo(app *app.App, name string) (*filemsg.FileInfoResponse, error) {
	fInf, err := app.GetFileInfo(name)
	if err != nil {
		s.logger.Error("Error occured during getting file info", zap.Error(err))
		return nil, err
	}

	return &filemsg.FileInfoResponse{Name: fInf.Name, Size: fInf.Size, Mode: uint32(fInf.Mode), Time: app.GetTimeAsInt64(fInf)}, nil
}

func (s *server) getFileNames(app *app.App) (*filemsg.FileListResponse, error) {
	names, err := app.GetFileNames()
	if err != nil {
		s.logger.Error("Error occured during getting file names", zap.Error(err))
		return nil, err
	}

	return &filemsg.FileListResponse{Name: names}, nil
}

// Если я правильно понимаю, то никакой особой проблемы с паматью/скоростью мы не должны получить
// так как беря слайсы я беру лишь ссылки на отрезки массива

func (s *server) getFileDividedData(app *app.App, name string) (*[][]byte, error) {
	f, err := app.GetFile(name)
	if err != nil {
		s.logger.Error("Error occured during getting file data", zap.Error(err))
		return nil, err
	}

	data := f.Data

	chData := make([][]byte, (len(data)+SendByteSize-1)/SendByteSize) // chunked data
	for i := 0; i < len(data); i += SendByteSize {
		chData[i/SendByteSize] = data[i:min(len(data), i+SendByteSize)]
	}

	return &chData, nil
}

func (s *server) sendData(stream filemsg.FileService_GetFileDataServer, data *[][]byte) error {
	for _, ch := range *data { // go through every data chunk
		toSend := &filemsg.FileDataResponse{Data: ch}
		if err := stream.Send(toSend); err != nil {
			s.logger.Fatal("Failed to send a message", zap.Error(err))
			return err
		}

		err := s.checkIfRecieved(stream)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server) checkIfRecieved(stream filemsg.FileService_GetFileDataServer) error {
	in, err := stream.Recv()
	if err == io.EOF {
		s.logger.Error("Unpredictable end of stream")
		return status.Errorf(codes.InvalidArgument, "Unpredictable end of stream")
	}
	if err != nil {
		s.logger.Fatal("Failed to receive request", zap.Error(err))
		return err
	}

	switch in.Param.(type) {
	case *filemsg.FileRequest_Got:
	default:
		s.logger.Error("Expected request with field Got")
		return status.Errorf(codes.InvalidArgument, "Expected request with field Got")
	}
	return nil
}
