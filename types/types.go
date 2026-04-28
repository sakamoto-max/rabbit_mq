package types

import (
	"time"
)

type Data struct {
	DbId          string         `db:"id"`
	TargetService string         `db:"target_service"`
	CreatedBy     string         `db:"created_by"`
	Task          string         `db:"task"`
	Status        string         `db:"status"`
	Payload       map[string]any `db:"payload"`
	CreatedAt     time.Time      `db:"created_at"`
	NumberOfTries int            `db:"number_of_tries"`
}
