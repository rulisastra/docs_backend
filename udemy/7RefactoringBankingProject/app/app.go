package app

import (
	"log"
	"net/http"
)

// jadi gunain mux (multiplexer) instead of http (default multiplexer)
func Start() {

	mux := http.NewServeMux()

	log.Println(`Starting server at http://localhost:8082`)

	mux.HandleFunc("/customers", getAllCustomer)

	// log.Fatal(http.ListenAndServe("localhost:8082", nil)) // addr dan handler (nil)
	log.Fatal(http.ListenAndServe("localhost:8082", mux))
}
