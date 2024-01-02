package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"homework/server/internal/file"
	"homework/server/internal/starboards/grpcclient"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientServerTestSuite struct {
	suite.Suite
	client    *grpcclient.Client
	addr      string
	fileData  [][]byte
	fileNames []string
	t         *testing.T
}

func (ts *ClientServerTestSuite) SetupSuite() {
	ts.client = CreateClient(ClientLogger{})
	ts.addr = ":8080"

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// create folder and files for file server repo

	fmt.Println(path)
	path += "/system/"
	os.MkdirAll(path, 0750)

	f1, err1 := os.Create(path + "123")
	f2, err2 := os.Create(path + "321")

	ts.fileNames = []string{"123", "321"}

	assert.NoError(ts.t, err1)
	assert.NoError(ts.t, err2)

	toWrite1 := make([]byte, 100)
	for i := 0; i < len(toWrite1); i++ {
		toWrite1[i] = byte(i)
	}

	toWrite2 := []byte{1, 2, 3, 4, 5}

	_, err1 = f1.Write(toWrite1)
	_, err2 = f2.Write(toWrite2)

	ts.fileData = [][]byte{toWrite1, toWrite2}

	assert.NoError(ts.t, err1)
	assert.NoError(ts.t, err2)

	err1 = f1.Close()
	err2 = f2.Close()

	assert.NoError(ts.t, err1)
	assert.NoError(ts.t, err2)

	go RunServer(ts.addr, path, ServerLoggers{})
}

func (ts *ClientServerTestSuite) TearDownSuite() {
	// don't forget to delete folder and files after all

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err1 := os.Remove(path + "/system/123")
	err2 := os.Remove(path + "/system/321")
	err3 := os.Remove(path + "/system")

	assert.NoError(ts.t, err1)
	assert.NoError(ts.t, err2)
	assert.NoError(ts.t, err3)
}

func createMap[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{})
	for _, x := range in {
		out[x] = struct{}{}
	}

	return out
}

func (ts *ClientServerTestSuite) TestNames() {
	names, err := ts.client.GetFileNames(ts.addr)
	assert.NoError(ts.t, err)

	res := createMap(names)
	exp := createMap(ts.fileNames)

	assert.Equal(ts.t, exp, res)
}

func (ts *ClientServerTestSuite) TestData() {
	data, err := ts.client.GetFileData(ts.addr, ts.fileNames[0])
	assert.NoError(ts.t, err)

	res := createMap(data)
	exp := createMap(ts.fileData[0])

	assert.Equal(ts.t, exp, res)
}

func (ts *ClientServerTestSuite) TestInfo() {
	info, err := ts.client.GetFileInfo(ts.addr, ts.fileNames[0])
	assert.NoError(ts.t, err)

	exp := file.FileInfo{Name: ts.fileNames[0], Size: info.Size, Mode: 420, ModTime: info.ModTime}

	assert.Less(ts.t, time.Since(info.ModTime), time.Second, "Too big difference between creation and now moments")
	assert.Equal(ts.t, exp, *info)
}

func TestServerHandler(t *testing.T) {
	ts := ClientServerTestSuite{t: t}
	suite.Run(t, &ts)
}
