package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Greeting page</h1>")
}

func main() {

	http.HandleFunc("/greet", greet)

	fmt.Println("Start server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
