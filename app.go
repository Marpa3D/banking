package main

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)

	fmt.Println("Start server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
