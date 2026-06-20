package store

import (
	"os"
	"path/filepath"
)

type ConfigAdapter interface {
	Read() ([]byte, error)
	Write(bytes []byte) error
}

type FileAdapter struct {
	Name string
}

var _ ConfigAdapter = (*FileAdapter)(nil)

func (a *FileAdapter) Read() ([]byte, error) {
	file, err := a.getPath()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(file)
}

func (a *FileAdapter) Write(bytes []byte) error {
	file, err := a.getPath()
	if err != nil {
		return err
	}
	return os.WriteFile(file, bytes, 0644)
}

func (a *FileAdapter) getPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, a.Name)
	return path, nil
}
