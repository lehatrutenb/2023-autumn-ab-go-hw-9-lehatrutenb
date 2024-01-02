package file

import (
	"os"
	"time"
)

// simular to os FileInfo interface, but without sys
type FileInfo struct {
	Name    string      // base name of the file
	Size    int64       // length in bytes for regular files; system-dependent for others
	Mode    os.FileMode // file mode bits
	ModTime time.Time   // modification time
}

type File struct {
	inf  os.FileInfo
	Data []byte
}

func NewFile(Inf os.FileInfo, Data []byte) *File {
	return &File{inf: Inf, Data: Data}
}

func (f *File) GetInfoData() FileInfo {
	return FileInfo{Name: f.inf.Name(), Size: f.inf.Size(),
		Mode: f.inf.Mode().Perm(), ModTime: f.inf.ModTime()}
}

func (f *FileInfo) GetTimeAsInt64() int64 {
	return f.ModTime.Unix()
}

// вообще, как будто бы хочется выносить подобные функции сюда, чтобы легче их менять
func GetInt64AsTime(x int64) time.Time {
	return time.Unix(x, 0)
}
