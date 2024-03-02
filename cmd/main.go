// cmd/main.go
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

	products := []product.Product{
		{Name: "Apple iPhone 15 Pro Max 256Gb серый", Price: 654857, Description: "Product Description 1", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hc1/h65/83559848181790.png?format=preview-large"},
		{Name: "Apple iPhone 13 128Gb Midnight черный", Price: 299929, Description: "Product Description 2", QuantityInStock: 3, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h32/h70/84378448199710.jpg?format=preview-large"},
		{Name: "Apple iPhone 15 128Gb черный", Price: 401794, Description: "Product Description 3", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/he2/h1d/83559338442782.png?format=preview-large"},
		{Name: "Apple iPhone 11 128Gb Slim Box черный", Price: 244900, Description: "Product Description 4", QuantityInStock: 2, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd8/hac/63897052413982.jpg?format=preview-large"},
		{Name: "Apple iPhone 13 128Gb белый", Price: 300000, Description: "Product Description 5", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hc9/h90/64209083007006.jpg?format=preview-large"},
		{Name: "Apple iPhone 14 128Gb черный", Price: 341178, Description: "Product Description 6", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h98/h2b/64400497737758.jpg?format=preview-large"},
		{Name: "Samsung Galaxy A24 6 ГБ/128 ГБ черный", Price: 91636, Description: "Product Description 7", QuantityInStock: 20, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hdc/h12/80750151303198.jpg?format=preview-large"},
		{Name: "Xiaomi Redmi 12 4G 8 ГБ/256 ГБ черный", Price: 68759, Description: "Product Description 8", QuantityInStock: 10, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h75/hbc/81335343775774.png?format=preview-large"},
		{Name: "Apple iPhone 15 Pro 256Gb синий", Price: 597788, Description: "Product Description 9", QuantityInStock: 30, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h88/hde/83559835697182.jpg?format=preview-large"},
		{Name: "Apple iPhone 14 128Gb starlight", Price: 341100, Description: "Product Description 10", QuantityInStock: 2, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h9f/h49/64481569832990.jpg?format=preview-large"},
		{Name: "Samsung Galaxy A54 5G 8 ГБ/256 ГБ черный", Price: 186756, Description: "Product Description 11", QuantityInStock: 11, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h81/h13/80435536887838.jpg?format=preview-large"},
		{Name: "Apple iPhone 13 128Gb зеленый", Price: 307856, Description: "Product Description 12", QuantityInStock: 10, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd1/h2f/64255724159006.jpg?format=preview-large"},
		{Name: "Apple iPhone 15 128Gb голубой", Price: 397652, Description: "Product Description 13", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd1/h07/83559339032606.png?format=preview-large"},
		{Name: "Xiaomi Redmi 12C 4 ГБ/128 ГБ серый", Price: 48800, Description: "Product Description 14", QuantityInStock: 16, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/ha6/h82/69586957697054.jpg?format=preview-large"},
		{Name: "Poco X6 Pro 12 ГБ/512 ГБ черный", Price: 172598, Description: "Product Description 15", QuantityInStock: 7, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hbe/h45/84940376899614.jpg?format=preview-large"},
		{Name: "Apple iPhone 15 128Gb черный", Price: 401794, Description: "Product Description 3", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/he2/h1d/83559338442782.png?format=preview-large"},
	}
	err = product.AddProductsToDB(db, products)
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
