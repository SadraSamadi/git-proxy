package store

import (
	"encoding/json"
)

type ConfigFormatter interface {
	Encode(config *Config) ([]byte, error)
	Decode(bytes []byte, config *Config) error
}

type JsonFormatter struct{}

var _ ConfigFormatter = (*JsonFormatter)(nil)

func (f *JsonFormatter) Encode(config *Config) ([]byte, error) {
	return json.MarshalIndent(config, "", "  ")
}

func (f *JsonFormatter) Decode(bytes []byte, config *Config) error {
	return json.Unmarshal(bytes, config)
}
