package store

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/yusuke0701/time-recorder/config"
)

var datastoreClient *datastore.Client

func init() {
	ctx := context.Background()

	if err := setNewDatastoreClient(ctx); err != nil {
		log.Fatal(err)
	}
}

func setNewDatastoreClient(ctx context.Context) (err error) {
	datastoreClient, err = datastore.NewClient(ctx, config.ProjectID)
	if err != nil {
		return fmt.Errorf("Failed to connect datastore: %v", err)
	}
	return nil
}
