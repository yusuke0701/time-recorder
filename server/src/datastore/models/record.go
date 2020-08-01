package models

import (
	"time"
)

type Record struct {
	ID    string    `json:"id" datastore:"-"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
