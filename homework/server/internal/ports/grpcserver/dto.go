package grpcserver

import (
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver/filemsg"
)

// TODO add optinos to change it
const SendByteSize = 1000

// Я не думаю, что прeобразвания в типы сообщений из protoc надо делать в app, так что думаю тут это хорошо

func GetFileInfo(app *app.App, name string) (*filemsg.FileInfoResponse, error) {
	fInf, err := app.GetFileInfo(name)
	if err != nil {
		return nil, err
	}

	return &filemsg.FileInfoResponse{Name: fInf.Name, Size: fInf.Size, Mode: uint32(fInf.Mode), Time: app.GetTimeAsInt64(fInf)}, nil
}

func GetFileNames(app *app.App) (*filemsg.FileListResponse, error) {
	names, err := app.GetFileNames()
	if err != nil {
		return nil, err
	}

	return &filemsg.FileListResponse{Name: names}, nil
}

// Если я правильно понимаю, то никакой особой проблемы с паматью/скоростью мы не должны получить
// так как беря слайсы я беру лишь ссылки на отрезки массива

func GetFileDividedData(app *app.App, name string) (*[][]byte, error) {
	f, err := app.GetFile(name)
	if err != nil {
		return nil, err
	}

	data := f.Data

	chData := make([][]byte, (len(data)+SendByteSize-1)/SendByteSize) // chunked data
	for i := 0; i < len(data); i += SendByteSize {
		chData[i/SendByteSize] = data[i:min(len(data), i+SendByteSize)]
	}

	return &chData, nil
}
