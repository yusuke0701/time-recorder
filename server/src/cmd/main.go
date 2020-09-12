package main

import (
	"log"
	"net/http"

	funcs "github.com/yusuke0701/time-recorder"
)

func main() {
	http.HandleFunc("/Records", funcs.Records)
	// http.HandleFunc("/GetGoogleID", funcs.GetGoogleID)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
