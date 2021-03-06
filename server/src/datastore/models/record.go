package models

import (
	"time"
)

type Record struct {
	ID          string    `json:"id" datastore:"-"`
	UserID      string    `json:"user_id"`
	Category    string    `json:"category"`
	Start       string    `json:"start"`
	StartDetail time.Time `json:"start_detail"`
	End         string    `json:"end"`
	EndDetail   time.Time `json:"end_detail"`
}
