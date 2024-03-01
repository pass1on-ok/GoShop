package main

import (
<<<<<<< HEAD
	"database/sql"
=======
>>>>>>> 935a73bd99255b86021e99277bece64d83106b06
	"html/template"
	"log"
	"net/http"

<<<<<<< HEAD
	"onlinestore/internal/handlers"
=======
>>>>>>> 935a73bd99255b86021e99277bece64d83106b06
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

<<<<<<< HEAD
	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProducts(w, r, db)
	}).Methods("GET")

	// Передаем переменную db в функцию homeHandler
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeHandler(w, r, db)
	}).Methods("GET")
=======
	r.HandleFunc("/", homeHandler).Methods("GET")
>>>>>>> 935a73bd99255b86021e99277bece64d83106b06

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

<<<<<<< HEAD
func homeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	products, err := product.GetAllProductsFromDB(db)
	if err != nil {
		http.Error(w, "Error retrieving products", http.StatusInternalServerError)
		return
	}
=======
func homeHandler(w http.ResponseWriter, r *http.Request) {

	products := product.GetAllProducts()
>>>>>>> 935a73bd99255b86021e99277bece64d83106b06

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
