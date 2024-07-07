package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
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
	w.Header().Add("Content-type", "Aplicattion/json")
	json.NewEncoder(w).Encode(customers)
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
