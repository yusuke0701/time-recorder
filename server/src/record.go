package funcs

import (
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

	category := r.FormValue("category")
	if category == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid request. category param is required.")
		return
	}

	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	googleID, err := callGetGoogleIDFunction(ctx, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// main process

	record := &models.Record{
		GoogleID:    googleID,
		Category:    category,
		StartDetail: timeutils.NowInJST(),
	}
	if err := (&store.Record{}).Upsert(ctx, record); err != nil {
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

	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	googleID, err := callGetGoogleIDFunction(ctx, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// main process

	record, err := (&store.Record{}).GetLastRecord(ctx, googleID)
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

	start := r.FormValue("start")
	if start != "" {
		if _, err := time.Parse(timeutils.DefaultFormat, start); start != "" && err != nil {
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

	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	googleID, err := callGetGoogleIDFunction(ctx, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// main process

	records, err := (&store.Record{}).List(ctx, googleID, start, end)
	if err == datastore.ErrNoSuchEntity {
		w.WriteHeader(http.StatusOK)
		return
	} else if err != nil {
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

	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	googleID, err := callGetGoogleIDFunction(ctx, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
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

	if record.GoogleID != googleID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	record.EndDetail = timeutils.NowInJST()

	if err := rStore.Upsert(ctx, record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
