package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux" // hadeh disuruh go get. ngga bisa apa langsungan
)

// https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/learn/lecture/22286162#overview

func Start() {

	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// define routes
	log.Println(`Starting server at http://localhost:8082`)
	log.Fatal(http.ListenAndServe("localhost:8082", router))
}
