package types

import (
	"encoding/json"
	"fmt"
)

// format for data when sending it between queues
type Data struct {
	DbId          string
	TaskName      string
	SentBy        string
	Payload       map[string]any
	TaskStatus    string
	TargetService string
	Err           error
}

func (d *Data) ConvertIntoBytes() (*[]byte, error) {

	dataInBytes, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("error in converting data into bytes : %w", err)
	}

	return &dataInBytes, nil
}
