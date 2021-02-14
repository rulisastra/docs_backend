package app

import (
	"2_ApplicationDevelopment/domain"
	"2_ApplicationDevelopment/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// define routes
	log.Println(`Starting server at http://localhost:8082`)
	log.Fatal(http.ListenAndServe("localhost:8082", router))
}
