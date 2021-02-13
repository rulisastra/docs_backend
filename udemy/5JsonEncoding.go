package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
gunain postman kuy
*/

type Customer struct {
	// tambahin json format kalo mau change the key in json response. ngga juga boleh
	Name string `json:"full_name" xml:"name"`
	City string
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ruli Sastra", "Palembang"}, // it should've been automatically, so we just need to input value. hmmm
		{"Kasmini", "Bone"},          // tapi ternyata 200 begini (tanpa deklarasi nama lagi)
	}

	// tambain ini kalo mau, ADD -> see the diffenrence
	w.Header().Add("Content-Type", "application/json")
	// ngga juga outputnya keluar "text/plain; charset=utf-8" di header postman

	json.NewEncoder(w).Encode(customers) // same with fprint
}

func main() {

	// liat di log
	log.Println(`Starting server at http://localhost:8082`)

	// define routes, post, put, dll
	http.HandleFunc("/customers", getAllCustomer)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8082", nil))

}
