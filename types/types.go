package types

import (
	"encoding/json"
	"fmt"
)

// format for data when sending it between queues
type Data struct {
	DbId          string         `json:"dbId"`
	TaskName      string         `json:"taskName"`
	SentBy        string         `json:"sentBy"`
	Payload       map[string]any `json:"payload"`
	TaskStatus    string         `json:"status"`
	TargetService string         `json:"targetService"`
	Err           error          `json:"err"`
}

func (d *Data) ConvertIntoBytes() (*[]byte, error) {

	dataInBytes, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("error in converting data into bytes : %w", err)
	}

	return &dataInBytes, nil
}
