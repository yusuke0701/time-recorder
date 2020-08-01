package store

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"

	"github.com/yusuke0701/goutils/manufacture"
	"github.com/yusuke0701/time-recorder/datastore/models"
)

type Record struct{}

func (r *Record) Get(ctx context.Context, id string) (record *models.Record, err error) {
	record = new(models.Record)
	if err := datastoreClient.Get(ctx, r.newKey(id), record); err != nil {
		return nil, fmt.Errorf("failed to get: %s", err)
	}

	return record, nil
}

func (r *Record) Upsert(ctx context.Context, record *models.Record) error {
	id, err := r.newID()
	if err != nil {
		return fmt.Errorf("failed to newID: %s", err)
	}

	if _, err := datastoreClient.Put(ctx, r.newKey(id), record); err != nil {
		return fmt.Errorf("failed to put: %s", err)
	}

	record.ID = id
	return nil
}

// for internal

func (*Record) kind() string {
	return "record"
}

func (*Record) newID() (string, error) {
	return manufacture.NewUUID()
}

func (r *Record) newKey(id string) *datastore.Key {
	return datastore.NameKey(r.kind(), id, nil)
}
