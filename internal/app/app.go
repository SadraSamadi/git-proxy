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
	config, err := store.Read()
	if err != nil {
		return err
	}
	config.Proxies[key] = value
	return store.Write(config)
}

func List() (map[string]string, error) {
	config, err := store.Read()
	if err != nil {
		return nil, err
	}
	return config.Proxies, nil
}

func Remove(key string) error {
	config, err := store.Read()
	if err != nil {
		return err
	}
	delete(config.Proxies, key)
	return store.Write(config)
}

func Use(key string) error {
	config, err := store.Read()
	if err != nil {
		return err
	}
	value, ok := config.Proxies[key]
	if !ok {
		return errors.New("key not found")
	}
	return git.Configure(value)
}

func Current() (string, error) {
	return git.Current()
}

func Unset() error {
	return git.Unset()
}
