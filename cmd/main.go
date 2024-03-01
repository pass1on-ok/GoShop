package main

import (
	"log"
	"net/http"

	"onlinestore/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", handlers.GetProducts).Methods("GET")

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
