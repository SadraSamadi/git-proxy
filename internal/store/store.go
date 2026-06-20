package store

import (
	"os"
)

var Formatter ConfigFormatter = &JsonFormatter{}
var Adapter ConfigAdapter = &FileAdapter{"config.json"}

func Read() (*Config, error) {
	config := DefaultConfig()
	bytes, err := Adapter.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return nil, err
	}
	err = Formatter.Decode(bytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func Write(config *Config) error {
	bytes, err := Formatter.Encode(config)
	if err != nil {
		return err
	}
	return Adapter.Write(bytes)
}
