package store

import (
	"encoding/json"
	"os"
)

var Adaptor DataAdapter = &JsonAdapter{"data.json"}

func Read() (*Data, error) {
	data := DefaultData()
	bytes, err := Adaptor.Read()
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
	return Adaptor.Write(bytes)
}
