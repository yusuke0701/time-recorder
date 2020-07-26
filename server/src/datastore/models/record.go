package models

import (
	"time"
)

type Record struct {
	ID    int64     `json:"id" datastore:"-"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
