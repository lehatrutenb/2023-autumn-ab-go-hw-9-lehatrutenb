package grpctest

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestFileInfo struct {
	FileData  [][]byte
	FileNames []string
}

// create folder and files for file server repo
func CreateTestFiles(t *testing.T) (TestFileInfo, string) {
	var fi TestFileInfo

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	path += "/system/"
	os.MkdirAll(path, 0750)

	f1, err1 := os.Create(path + "123")
	f2, err2 := os.Create(path + "321")

	fi.FileNames = []string{"123", "321"}

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	toWrite1 := make([]byte, 100)
	for i := 0; i < len(toWrite1); i++ {
		toWrite1[i] = byte(i)
	}

	toWrite2 := []byte{1, 2, 3, 4, 5}

	_, err1 = f1.Write(toWrite1)
	_, err2 = f2.Write(toWrite2)

	fi.FileData = [][]byte{toWrite1, toWrite2}

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	err1 = f1.Close()
	err2 = f2.Close()

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	return fi, path
}

// don't forget to delete folder and files after test
func RemoveTestFiles(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err1 := os.Remove(path + "/system/123")
	err2 := os.Remove(path + "/system/321")
	err3 := os.Remove(path + "/system")

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NoError(t, err3)
}

func CreateMap[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{})
	for _, x := range in {
		out[x] = struct{}{}
	}

	return out
}
