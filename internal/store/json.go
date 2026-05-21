package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (*Data, error) {
	file, err := getPath()
	if err != nil {
		return nil, err
	}
	bytes, err := os.ReadFile(file)
	data := getDefault()
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return nil, err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Write(data *Data) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	file, err := getPath()
	if err != nil {
		return err
	}
	return os.WriteFile(file, bytes, 0644)
}

func getPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "data.json")
	return path, nil
}
