package foldersys

import (
	"homework/server/internal/file"
	"homework/server/internal/grpctest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type AppTestSuite struct {
	suite.Suite
	grpctest.TestFileInfo
	sys *RepositorySys
	t   *testing.T
}

func (ts *AppTestSuite) SetupSuite() {
	fi, path := grpctest.CreateTestFiles(ts.t)
	ts.FileNames, ts.FileData = fi.FileNames, fi.FileData

	logger, err := zap.NewProduction(zap.WithCaller(true))
	assert.NoError(ts.t, err)

	ts.sys = NewRepo(path, logger)
}

func (ts *AppTestSuite) TearDownSuite() {
	grpctest.RemoveTestFiles(ts.t)
}

func (ts *AppTestSuite) TestGetFileNames() {
	names, err := ts.sys.GetFileNames()
	assert.NoError(ts.t, err)

	res := grpctest.CreateMap(names)
	exp := grpctest.CreateMap(ts.FileNames)

	assert.Equal(ts.t, exp, res)
}

func (ts *AppTestSuite) TestGetFileByNameData() {
	f, err := ts.sys.GetFileByName(ts.FileNames[0])
	assert.NoError(ts.t, err)

	res := grpctest.CreateMap(f.Data)
	exp := grpctest.CreateMap(ts.FileData[0])

	assert.Equal(ts.t, exp, res)
}

func (ts *AppTestSuite) TestGetFileByNameInfo() {
	f, err := ts.sys.GetFileByName(ts.FileNames[0])
	assert.NoError(ts.t, err)

	info := f.GetInfoData()

	exp := file.FileInfo{Name: ts.FileNames[0], Size: info.Size, Mode: 420, ModTime: info.ModTime}

	assert.Less(ts.t, time.Since(info.ModTime), time.Second, "Too big difference between creation and now moments")
	assert.Equal(ts.t, exp, info)
}

func TestRepoMainFuncs(t *testing.T) {
	ts := AppTestSuite{t: t}
	suite.Run(t, &ts)
}
