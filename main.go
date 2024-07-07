package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Viacheslav",
			"Zhukovsky",
			"18185",
		},
		{"Mark",
			"Moskow",
			"147800",
		},
	}

	if r.Header.Get("Content-Type") == "Aplication/xml" {

		w.Header().Add("Content-Type", "Aplicattion/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "Aplication/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Greeting page</h1>")
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)

	fmt.Println("Start server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
