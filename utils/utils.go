package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sakamoto-max/rabbit_mq/types"
)

func ConvertIntoBytes(payload any) (*[]byte, error) {

	dataInBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error in converting data into bytes : %w", err)
	}

	return &dataInBytes, nil
}

func ConvertIntoJosn(data *[]byte) *types.Data {

	var D types.Data

	_ = json.Unmarshal(*data, &D)

	return &D
}
