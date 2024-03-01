package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pass1on-ok/Golang-Project/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	http.ListenAndServe(":8080", r)
}
