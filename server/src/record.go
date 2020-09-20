package funcs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/yusuke0701/time-recorder/datastore/models"
	"github.com/yusuke0701/time-recorder/datastore/store"
	timeutils "github.com/yusuke0701/time-recorder/time"
)

// Records は、RecordAPIの関数です。
func Records(w http.ResponseWriter, r *http.Request) {
	// pre process

	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	userID, err := callVerifyIDTokenFunction(ctx, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// main process

	switch r.Method {
	case http.MethodPost:
		createRecord(ctx, w, r, userID)
	case http.MethodGet:
		last := r.FormValue("last")
		if last != "" {
			getLastRecord(ctx, w, r, userID)
		} else {
			listRecord(ctx, w, r, userID)
		}
	case http.MethodPut:
		updateRecord(ctx, w, r, userID)
	case http.MethodDelete:
		// TODO:
		w.WriteHeader(http.StatusMethodNotAllowed)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

func createRecord(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	category := r.FormValue("category")
	if category == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid request. category param is required.")
		return
	}

	record := &models.Record{
		UserID:      userID,
		Category:    category,
		StartDetail: timeutils.NowInJST(),
	}
	if err := (&store.Record{}).Upsert(ctx, record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}

func getLastRecord(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	record, err := (&store.Record{}).GetLastRecord(ctx, userID)
	if err == datastore.ErrNoSuchEntity {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}

func listRecord(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	start := r.FormValue("start")
	if start != "" {
		if _, err := time.Parse(timeutils.DefaultFormat, start); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
			return
		}
	}

	end := r.FormValue("end")
	if end != "" {
		if _, err := time.Parse(timeutils.DefaultFormat, end); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
			return
		}
	}

	records, err := (&store.Record{}).List(ctx, userID, start, end)
	if err == datastore.ErrNoSuchEntity {
		w.WriteHeader(http.StatusOK)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(records)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}

func updateRecord(ctx context.Context, w http.ResponseWriter, r *http.Request, userID string) {
	var recordID string
	{
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}
		defer r.Body.Close()

		recordID = string(body)
		if recordID == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Set record id at body.")
			return
		}
	}

	record, err := (&store.Record{}).Get(ctx, recordID)
	if err == datastore.ErrNoSuchEntity {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	if record.UserID != userID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	record.EndDetail = timeutils.NowInJST()

	if err := (&store.Record{}).Upsert(ctx, record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(b))
}
