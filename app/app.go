package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customers/{customers_id:[0-9]+}", getCustomer)

	fmt.Println("Start server...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customers_id"])
}
