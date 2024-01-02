package app

import (
	"homework/server/internal/file"
	"time"
)

type App struct {
	repo file.Repository
}

func NewApp(repo file.Repository) App {
	return App{repo: repo}
}

func (a *App) GetFile(Name string) (*file.File, error) {
	f, err := a.repo.GetFileByName(Name)
	if err != nil {
		return nil, err
	}

	return f, err
}

// в плане оптимизации это плохое решение, но в плане читаемости, мне кажется хорошее

func (a *App) GetFileInfo(Name string) (*file.FileInfo, error) {
	f, err := a.repo.GetFileByName(Name)
	if err != nil {
		return nil, err
	}

	return ptr(f.GetInfoData()), err
}

func ptr[T any](x T) *T {
	return &x
}

func (a *App) GetFileNames() ([]string, error) {
	return a.repo.GetFileNames()
}

func (*App) GetTimeAsInt64(f *file.FileInfo) int64 {
	return f.GetTimeAsInt64()
}

func (*App) GetInt64AsTime(x int64) time.Time {
	return file.GetInt64AsTime(x)
}
