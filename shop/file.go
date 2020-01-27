package shop

import (
	"bytes"
	"encoding/json"
)

func (shop Shop) Import(data []byte) error {
	if err := json.Unmarshal(data, &shop); err != nil {
		return err
	}
	return nil
}

func (shop Shop) Export() ([]byte, error) {
	reqBodyBytes := new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBytes).Encode(shop); err != nil {
		return nil, err
	}

	return reqBodyBytes.Bytes(), nil
}

func (shop Shop) WriteToFile(file string, data []byte) error {
	return nil
}
