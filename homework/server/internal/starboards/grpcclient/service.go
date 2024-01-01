package grpcclient

import (
	"context"
	"homework/server/internal/file"
	"homework/server/internal/ports/grpcserver/filemsg"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO add option settings for it

const RequestTimeout = time.Second
const StreamTimeout = time.Second * 10

func GetFileInfo(addr string, fileName string) (*file.FileInfo, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := filemsg.NewFileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	fi, err := c.GetFileInfo(ctx, &filemsg.FileRequest{Param: &filemsg.FileRequest_Name{Name: fileName}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return nil, err
	}

	return GrpcFileInfoToServerFileInfo(fi), nil
}

func GetFileNames(addr string) ([]string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := filemsg.NewFileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	names, err := c.GetFileNames(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return nil, err
	}

	return names.Name, nil
}

func GetFileData(addr string, fileName string) ([]byte, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := filemsg.NewFileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), StreamTimeout)
	defer cancel()

	stream, err := c.GetFileData(ctx)
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
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

	return data, err
}

func sendFileName(stream filemsg.FileService_GetFileDataClient, fn string) error {
	err := stream.Send(&filemsg.FileRequest{Param: &filemsg.FileRequest_Name{Name: fn}})
	return err
}

func getFileData(stream filemsg.FileService_GetFileDataClient) ([]byte, error) {
	data := newFileData()

	for {
		in, err := stream.Recv()
		if err == io.EOF { // End of stream
			break
		}
		if err != nil {
			log.Printf("Failed to receive a message : %v", err)
			return []byte{}, err
		}

		data.addData(in.Data)
		stream.Send(&filemsg.FileRequest{Param: &filemsg.FileRequest_Got{Got: true}})
	}

	return data.getData(), nil
}
