package store

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/datastore"

	"github.com/yusuke0701/goutils/manufacture"
	"github.com/yusuke0701/time-recorder/datastore/models"
	timeutils "github.com/yusuke0701/time-recorder/time"
)

type Record struct{}

func (r *Record) Get(ctx context.Context, id string) (*models.Record, error) {
	record := new(models.Record)
	if err := datastoreClient.Get(ctx, r.newKey(id), record); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get a record: %s", err)
	}

	record.ID = id
	record.StartDetail = timeutils.InJST(record.StartDetail)
	record.EndDetail = timeutils.InJST(record.EndDetail)
	return record, nil
}

func (r *Record) GetLastRecord(ctx context.Context, googleID string) (*models.Record, error) {
	q := datastore.NewQuery(r.kind())
	q = q.Filter("GoogleID =", googleID)
	q = q.Filter("End =", "0001-01-1")

	var records []*models.Record
	keys, err := datastoreClient.GetAll(ctx, q, &records)
	if err != nil {
		return nil, fmt.Errorf("failed to list record: %s", err)
	}
	if len(records) == 0 {
		return nil, datastore.ErrNoSuchEntity
	}
	if len(records) > 1 {
		return nil, errors.New("TODO: 管理者対応")
	}

	r.setID(keys, records)
	for _, record := range records {
		record.StartDetail = timeutils.InJST(record.StartDetail)
		record.EndDetail = timeutils.InJST(record.EndDetail)
	}
	return records[0], nil
}

func (r *Record) List(ctx context.Context, googleID, start, end string) (records []*models.Record, err error) {
	q := datastore.NewQuery(r.kind())
	{
		q = q.Filter("GoogleID =", googleID)
		if start != "" {
			q = q.Filter("Start =", start)
		}
		if end != "" {
			q = q.Filter("End =", end)
		}
		q = q.Order("StartDetail")
	}

	keys, err := datastoreClient.GetAll(ctx, q, &records)
	if err != nil {
		return nil, fmt.Errorf("failed to list record: %s", err)
	}
	if len(records) == 0 {
		return nil, datastore.ErrNoSuchEntity
	}

	r.setID(keys, records)
	for _, record := range records {
		record.StartDetail = timeutils.InJST(record.StartDetail)
		record.EndDetail = timeutils.InJST(record.EndDetail)
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
	record.Start = record.StartDetail.Format(timeutils.DefaultFormat)
	record.End = record.EndDetail.Format(timeutils.DefaultFormat)

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

func (*Record) setID(keys []*datastore.Key, records []*models.Record) {
	for i := range keys {
		records[i].ID = keys[i].Name
	}
}
