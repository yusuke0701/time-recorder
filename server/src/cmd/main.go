package main

import (
	"log"
	"net/http"

	funcs "github.com/yusuke0701/time-recorder"
)

func main() {
	http.HandleFunc("/CreateRecord", funcs.CreateRecord)
	http.HandleFunc("/GetLastRecord", funcs.GetLastRecord)
	http.HandleFunc("/ListRecord", funcs.ListRecord)
	http.HandleFunc("/UpdateRecord", funcs.UpdateRecord)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
