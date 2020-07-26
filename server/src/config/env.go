package config

import (
	"errors"
	"log"
	"os"
)

var (
	ProjectID string
)

func init() {
	if err := loadFromEnv(); err != nil {
		log.Fatal(err)
	}
}

func loadFromEnv() error {
	ProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if ProjectID == "" {
		return errors.New("failed init. Please set GOOGLE_CLOUD_PROJECT")
	}
	return nil
}
