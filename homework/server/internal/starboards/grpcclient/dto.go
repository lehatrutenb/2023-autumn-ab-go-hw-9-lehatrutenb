package grpcclient

import (
	"homework/server/internal/file"
	"homework/server/internal/ports/grpcserver/filemsg"
	"io/fs"
	"time"
)

func GrpcFileInfoToServerFileInfo(fi *filemsg.FileInfoResponse) *file.FileInfo {
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
