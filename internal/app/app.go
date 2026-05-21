package app

import (
	"errors"

	"github.com/SadraSamadi/git-proxy/internal/git"
	"github.com/SadraSamadi/git-proxy/internal/store"
)

func Save(key, value string) error {
	if !store.Validate(value) {
		return errors.New("invalid value")
	}
	data, err := store.Read()
	if err != nil {
		return err
	}
	data.Proxies[key] = value
	return store.Write(data)
}

func List() (map[string]string, error) {
	data, err := store.Read()
	if err != nil {
		return nil, err
	}
	return data.Proxies, nil
}

func Remove(key string) error {
	data, err := store.Read()
	if err != nil {
		return err
	}
	delete(data.Proxies, key)
	return store.Write(data)
}

func Use(key string) error {
	data, err := store.Read()
	if err != nil {
		return err
	}
	value, ok := data.Proxies[key]
	if !ok {
		return errors.New("key not found")
	}
	return git.Configure(value)
}

func Unset() error {
	return git.Unset()
}

func Current() (string, error) {
	return git.Current()
}
