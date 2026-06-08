package store

import (
	"os"
	"path/filepath"
)

type ConfigAdapter interface {
	Read() ([]byte, error)
	Write(bytes []byte) error
}

type JsonAdapter struct {
	Name string
}

var _ ConfigAdapter = (*JsonAdapter)(nil)

func (a *JsonAdapter) Read() ([]byte, error) {
	file, err := a.getPath()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(file)
}

func (a *JsonAdapter) Write(bytes []byte) error {
	file, err := a.getPath()
	if err != nil {
		return err
	}
	return os.WriteFile(file, bytes, 0644)
}

func (a *JsonAdapter) getPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, a.Name)
	return path, nil
}
