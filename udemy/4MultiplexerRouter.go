package main

import (
	"fmt"
	"log"
	"net/http"
)

func mulai(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat! karena sudah memulai hal-hal kecil!")
}

// main
func print() {

	// define routes, post, put, dll
	http.HandleFunc("/mulai", mulai)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8082", nil))

	// liat di log
	fmt.Print("Starting server at http://localhost:8082")
}
