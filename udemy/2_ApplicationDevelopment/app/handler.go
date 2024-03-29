package app

import (
	"2_ApplicationDevelopment/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"` // penamaan yang konstan itu perlu!
}

// membuat implementasi konkret, "CustomerHandlers struct"
// yang punya dependencies service
type CustomerHandlers struct {
	service service.CustomerService // sebagai dependencies dari CustomerHandlers (handlers)
	// kenapa ngga bisa export? huhuhu
}

// getAllCustomers
// ch adalah the right receiver
// yang pointing to CustomerHandlers
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Ruli Sastra", "Palembang"},
	// 	{"Kasmini", "Bone"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	// jika request header -> xml
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

// map of all the segment
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
