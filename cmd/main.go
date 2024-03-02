package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
						id SERIAL PRIMARY KEY,
						name TEXT,
						price NUMERIC,
						description TEXT,
						quantity_in_stock INTEGER,
						imagepath TEXT
					)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO products (name, price, description, quantity_in_stock, imagepath)
						VALUES ($1, $2, $3, $4, $5)`,
		"Product Name", 19.99, "Product Description", 100, "https://resources.cdn-kaspi.kz/img/m/p/hc1/h65/83559848181790.png?format=preview-large")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data inserted successfully")

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("web-page"))
	http.Handle("/web-page/", http.StripPrefix("/web-page/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeHandler(w, r, db)
	}).Methods("GET")

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	products, err := product.GetAllProductsFromDB(db)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error retrieving products:", err)
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
