package types

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	DbId          string
	TaskName      string
	Payload       map[string]any
	TaskStatus    string
	SentBy        string
	TargetService string
}

func (d *Data) ConvertIntoBytes() (*[]byte, error) {

	dataInBytes, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("error in converting data into bytes : %w", err)
	}

	return &dataInBytes, nil
}