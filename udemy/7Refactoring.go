package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

/*
gunain postman kuy
*/

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ruli Sastra", "Palembang"},
		{"Kasmini", "Bone"},
	}

	// jika request header -> xml
	if r.Header.Get("Content-Type") == "application/xml" { // PERHATIKAN booleannya
		w.Header().Add("Content-Type", "application/xml") // tulis header dan add xml
		xml.NewEncoder(w).Encode(customers)               // print xml
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func main() {

	log.Println(`Starting server at http://localhost:8082`)

	http.HandleFunc("/customers", getAllCustomer)

	log.Fatal(http.ListenAndServe("localhost:8082", nil))

}