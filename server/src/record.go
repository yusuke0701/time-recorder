package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/yusuke0701/time-recorder/datastore/models"
	"github.com/yusuke0701/time-recorder/datastore/store"
	"github.com/yusuke0701/time-recorder/time"
)

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	// pre process

	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// main process

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

func GetRecord(w http.ResponseWriter, r *http.Request) {
	// pre process

	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
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

	// main process

	record, err := (&store.Record{}).Get(ctx, id)
	if err == datastore.ErrNoSuchEntity {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, record.ID)
}

func GetLastRecord(w http.ResponseWriter, r *http.Request) {
	// pre process

	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// main process

	record, err := (&store.Record{}).GetLastRecord(ctx)
	if err == datastore.ErrNoSuchEntity {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, record.ID)
}

func ListRecord(w http.ResponseWriter, r *http.Request) {
	// pre process

	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// main process

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

func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	// pre process

	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		setHeaderForCORS(w)
		return
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
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

	// main process

	var rStore *store.Record

	record, err := rStore.Get(ctx, id)
	if err == datastore.ErrNoSuchEntity {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
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
