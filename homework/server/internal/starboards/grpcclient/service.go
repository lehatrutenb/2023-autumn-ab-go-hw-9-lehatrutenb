package grpcclient

import (
	"context"
	"homework/server/internal/file"
	"homework/server/internal/ports/grpcserver/filemsg"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO add option settings for it

const (
	RequestTimeout   = time.Second
	StreamTimeout    = time.Second * 10
	MaxLoggedDataLen = 15
)

func (c *client) GetFileInfo(addr string, fileName string) (*file.FileInfo, error) {
	c.logger.Info("Send get file info request",
		zap.String("addr", addr), zap.String("fileName", fileName))

	fsc, conn, err := c.createNewFileServiceClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	fi, err := fsc.GetFileInfo(ctx, &filemsg.FileRequest{Param: &filemsg.FileRequest_Name{Name: fileName}})
	if err != nil {
		c.logger.Error("Failed to get file info", zap.Error(err))
		return nil, err
	}

	c.logger.Info("Get file info response",
		zap.String("Name", fi.Name), zap.Int64("Size", fi.Size), zap.Uint32("Mode", fi.Mode), zap.Int64("ModTime", fi.Time))

	return grpcFileInfoToServerFileInfo(fi), nil
}

func (c *client) GetFileNames(addr string) ([]string, error) {
	c.logger.Info("Send get file names request", zap.String("addr", addr))

	fsc, conn, err := c.createNewFileServiceClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	names, err := fsc.GetFileNames(ctx, &emptypb.Empty{})
	if err != nil {
		c.logger.Error("Failed to get file names", zap.Error(err))
		return nil, err
	}

	c.logger.Info("Get file names response",
		zap.Strings("Names", names.Name))

	return names.Name, nil
}

func (c *client) GetFileData(addr string, fileName string) ([]byte, error) {
	c.logger.Info("Send get file data request",
		zap.String("addr", addr), zap.String("fileName", fileName))

	fsc, conn, err := c.createNewFileServiceClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), StreamTimeout)
	defer cancel()

	stream, err := fsc.GetFileData(ctx)
	if err != nil {
		c.logger.Error("Failed to get file data", zap.Error(err))
		return nil, err
	}

	err = sendFileName(stream, fileName)
	if err != nil {
		return nil, err
	}

	data, err := getFileData(stream)
	if err != nil {
		return nil, err
	}

	if len(data) <= MaxLoggedDataLen {
		c.logger.Info("Get data response",
			zap.Int("Size", len(data)), zap.Binary("Data", data))
	} else {
		c.logger.Info("Get data response",
			zap.Int("Size", len(data)))
	}

	return data, err
}
