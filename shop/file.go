package shop

import (
	"bytes"
	"encoding/json"
)

func (shop Shop) Import(data []byte) error {
	return json.Unmarshal(data, &shop)
}

func (shop Shop) Export() ([]byte, error) {
	reqBodyBytes := new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBytes).Encode(shop); err != nil {
		return nil, err
	}

	return reqBodyBytes.Bytes(), nil
}

//func (shop Shop) WriteToFile(file string, data []byte) error {
//	return nil
//}
