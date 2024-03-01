package main

import (
	"html/template"
	"log"
	"net/http"

	"onlinestore/pkg/product"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("web-page"))
	http.Handle("/web-page/", http.StripPrefix("/web-page/", fs))

	r.HandleFunc("/", homeHandler).Methods("GET")

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	products := product.GetAllProducts()

	tmpl, err := template.ParseFiles("web-page/homepage/home.html")
	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
		return
	}
}
