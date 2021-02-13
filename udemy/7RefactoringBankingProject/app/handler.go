package app

// handlers berisi seluruh FUNGSI yang bertugas menyerahkan tugas dari DATABASE - MAIN

import (
	"encoding/json"
	"encoding/xml"
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
