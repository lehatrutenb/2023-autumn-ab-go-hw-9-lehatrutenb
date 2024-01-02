package foldersys

import (
	"homework/server/internal/file"
	"os"
	"sync"

	"go.uber.org/zap"
)

type RepositorySys struct {
	mu      sync.Mutex
	dirPath string
	logger  *zap.Logger
}

func NewRepo(dirPath string, lr *zap.Logger) *RepositorySys {
	return &RepositorySys{mu: sync.Mutex{}, dirPath: dirPath, logger: lr}
}

// TODO logic
func (repo *RepositorySys) AddFile(f file.File) error {
	repo.logger.Fatal("Tried to use unimplemented AddFile")
	return nil
}

func (repo *RepositorySys) GetFileByName(Name string) (*file.File, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	fi, err := os.Stat(repo.dirPath + "/" + Name)
	if err != nil {
		repo.logger.Error("Failed to get file info", zap.Error(err))
		return nil, err
	}

	data, err := os.ReadFile(repo.dirPath + "/" + Name)
	if err != nil {
		repo.logger.Error("Failed to read file", zap.Error(err))
		return nil, err
	}

	return file.NewFile(fi, data), nil
}

func (repo *RepositorySys) GetFileNames() ([]string, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	files, err := os.ReadDir(repo.dirPath)
	if err != nil {
		repo.logger.Error("Failed to get dir info", zap.Error(err))
		return nil, err
	}

	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Name()
	}

	return names, nil
}
