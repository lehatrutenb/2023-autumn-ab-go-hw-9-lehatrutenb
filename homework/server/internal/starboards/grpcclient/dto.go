package grpcclient

import (
	"homework/server/internal/file"
	"homework/server/internal/ports/grpcserver/filemsg"
	"io"
	"io/fs"
	"log"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// dto funcs promise to log errors themselve

func (c *Client) createNewFileServiceClient(addr string, opts ...grpc.DialOption) (filemsg.FileServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		c.logger.Fatal("Failed to connect to server", zap.Error(err))
		return nil, nil, err
	}

	return filemsg.NewFileServiceClient(conn), conn, nil
}

func grpcFileInfoToServerFileInfo(fi *filemsg.FileInfoResponse) *file.FileInfo {
	return &file.FileInfo{Name: fi.Name, Size: fi.Size, Mode: fs.FileMode(fi.Mode), ModTime: time.Unix(fi.Time, 0)}
}

type fileData []byte

func newFileData() fileData {
	return make([]byte, 0)
}

func (fd *fileData) addData(data []byte) {
	*fd = append(*fd, data...)
}

func (fd *fileData) getData() []byte {
	return []byte(*fd)
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
