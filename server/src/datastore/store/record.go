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
		return nil, fmt.Errorf("failed to get a record: %s", err)
	}

	record.ID = id

	return record, nil
}

func (r *Record) List(ctx context.Context) (records []*models.Record, err error) {
	q := datastore.NewQuery(r.kind())

	if _, err := datastoreClient.GetAll(ctx, q, &records); err != nil {
		return nil, fmt.Errorf("failed to list record: %s", err)
	}

	return records, nil
}

func (r *Record) Upsert(ctx context.Context, record *models.Record) error {
	if record.ID == "" {
		id, err := r.newID()
		if err != nil {
			return fmt.Errorf("failed to create a new record ID: %s", err)
		}
		record.ID = id
	}

	if _, err := datastoreClient.Put(ctx, r.newKey(record.ID), record); err != nil {
		return fmt.Errorf("failed to put a record: %s", err)
	}

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
