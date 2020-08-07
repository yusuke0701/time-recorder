package models

import (
	"time"
)

type Record struct {
	ID       string    `json:"id" datastore:"-"`
	GoogleID string    `json:"google_id"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}
