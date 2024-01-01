package foldersys

import (
	"homework/server/internal/file"
	"os"
	"sync"
)

type RepositorySys struct {
	mu      sync.Mutex
	dirPath string
}

func NewRepo(dirPath string) *RepositorySys {
	return &RepositorySys{mu: sync.Mutex{}, dirPath: dirPath}
}

// TODO logic
func (repo *RepositorySys) AddFile(f file.File) error {
	return nil
}

func (repo *RepositorySys) GetFileByName(Name string) (*file.File, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	fi, err := os.Stat(repo.dirPath + "/" + Name)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(repo.dirPath + "/" + Name)
	if err != nil {
		return nil, err
	}

	return file.NewFile(fi, data), nil
}

func (repo *RepositorySys) GetFileNames() ([]string, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	files, err := os.ReadDir(repo.dirPath)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Name()
	}

	return names, nil
}
