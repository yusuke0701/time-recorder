package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/yusuke0701/time-recorder/datastore/models"
	"github.com/yusuke0701/time-recorder/datastore/store"
	"github.com/yusuke0701/time-recorder/datastore/time"
)

func Start(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	record := &models.Record{
		Start: time.NowInJST(),
	}

	if err := (&store.Record{}).Upsert(ctx, record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, record.ID)
}

func End(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}
	defer r.Body.Close()
	id := string(body)

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Set id in body.")
		return
	}

	var rStore *store.Record

	record, err := rStore.Get(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	record.End = time.NowInJST()

	if err := rStore.Upsert(ctx, record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func GetLastRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	record, err := (&store.Record{}).GetLastRecord(ctx)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, record.ID)
}

func ListRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	records, err := (&store.Record{}).List(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	recordsBytes, err := json.Marshal(records)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(recordsBytes))
}
