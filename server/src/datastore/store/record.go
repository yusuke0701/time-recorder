package store

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/yusuke0701/time-recorder/datastore/models"
)

type Record struct{}

func (r *Record) Get(ctx context.Context, id int64) (record *models.Record, err error) {
	if err := datastoreClient.Get(ctx, r.idKey(id), record); err != nil {
		return nil, err
	}

	return record, nil
}

func (r *Record) Upsert(ctx context.Context, record *models.Record) error {
	key, err := datastoreClient.Put(ctx, r.incompleteKey(), record)
	if err != nil {
		return err
	}
	record.ID = key.ID

	return nil
}

// for internal

func (*Record) kind() string {
	return "record"
}

func (r *Record) incompleteKey() *datastore.Key {
	return datastore.IncompleteKey(r.kind(), nil)
}

func (r *Record) idKey(id int64) *datastore.Key {
	return datastore.IDKey(r.kind(), id, nil)
}
