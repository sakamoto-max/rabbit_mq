package utils

import (
	"encoding/json"
	"github.com/sakamoto-max/rabbit_mq/types"
)

func ConvertIntoJson(data *[]byte) *types.Data {

	var D types.Data

	_ = json.Unmarshal(*data, &D)

	return &D
}
