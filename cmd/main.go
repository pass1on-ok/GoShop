package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"onlinestore/internal/handlers"
	"onlinestore/pkg/product"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:kumar@localhost/Online%20Store?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("web-page"))
	http.Handle("/web-page/", http.StripPrefix("/web-page/", fs))

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProducts(w, r, db)
	}).Methods("GET")

	// Передаем переменную db в функцию homeHandler
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeHandler(w, r, db)
	}).Methods("GET")

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	products, err := product.GetAllProductsFromDB(db)
	if err != nil {
		http.Error(w, "Error retrieving products", http.StatusInternalServerError)
		return
	}

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
