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
