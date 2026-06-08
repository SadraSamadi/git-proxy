package store

import (
	"encoding/json"
	"os"
)

var Adapter ConfigAdapter = &JsonAdapter{"config.json"}

func Read() (*Config, error) {
	config := DefaultConfig()
	bytes, err := Adapter.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return nil, err
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func Write(config *Config) error {
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return Adapter.Write(bytes)
}
