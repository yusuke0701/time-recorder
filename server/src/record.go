package funcs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yusuke0701/time-recorder/datastore/models"
	"github.com/yusuke0701/time-recorder/datastore/store"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func Start(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	record := &models.Record{
		Start: time.Now().In(jst),
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}
	defer r.Body.Close()
	id := string(body)

	var rStore *store.Record

	record, err := rStore.Get(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	record.End = time.Now().In(jst)

	if err := rStore.Upsert(ctx, record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
