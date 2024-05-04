// cmd/main.go
package main

//docker run -p 8080:8080 --network=my-network my-golang-app
import (
	"database/sql"
	"log"
	"net/http"
	"onlinestore/internal/handlers"
	"onlinestore/pkg/cart"
	"onlinestore/pkg/product"
	"onlinestore/pkg/user"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:kumar@localhost/Online%20Store?sslmode=disable")
	//db, err := sql.Open("postgres", "postgres://postgres:kumar@my-postgres/Online%20Store?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = product.EnsureTableExists(db)
	if err != nil {
		log.Fatal(err)
	}
	products := []product.Product{
		{Name: "Apple iPhone 15 Pro Max 256Gb Gray", Price: 654857, Description: "Product Description 1", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hc1/h65/83559848181790.png?format=preview-large"},
		{Name: "Apple iPhone 13 128Gb Midnight Black", Price: 299929, Description: "Product Description 2", QuantityInStock: 3, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h32/h70/84378448199710.jpg?format=preview-large"},
		{Name: "Apple iPhone 15 128Gb Black", Price: 401794, Description: "Product Description 3", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/he2/h1d/83559338442782.png?format=preview-large"},
		{Name: "Apple iPhone 11 128Gb Slim Box Black", Price: 244900, Description: "Product Description 4", QuantityInStock: 2, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd8/hac/63897052413982.jpg?format=preview-large"},
		{Name: "Apple iPhone 13 128Gb White", Price: 300000, Description: "Product Description 5", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hc9/h90/64209083007006.jpg?format=preview-large"},
		{Name: "Apple iPhone 14 128Gb Black", Price: 341178, Description: "Product Description 6", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h98/h2b/64400497737758.jpg?format=preview-large"},
		{Name: "Samsung Galaxy A24 6GB/128GB Black", Price: 91636, Description: "Product Description 7", QuantityInStock: 20, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hdc/h12/80750151303198.jpg?format=preview-large"},
		{Name: "Xiaomi Redmi 12 4G 8GB/256GB Black", Price: 68759, Description: "Product Description 8", QuantityInStock: 10, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h75/hbc/81335343775774.png?format=preview-large"},
		{Name: "Apple iPhone 15 Pro 256Gb Blue", Price: 597788, Description: "Product Description 9", QuantityInStock: 30, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h88/hde/83559835697182.jpg?format=preview-large"},
		{Name: "Apple iPhone 14 128Gb Starlight", Price: 341100, Description: "Product Description 10", QuantityInStock: 2, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h9f/h49/64481569832990.jpg?format=preview-large"},
		{Name: "Samsung Galaxy A54 5G 8GB/256GB Black", Price: 186756, Description: "Product Description 11", QuantityInStock: 11, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h81/h13/80435536887838.jpg?format=preview-large"},
		{Name: "Apple iPhone 13 128Gb Green", Price: 307856, Description: "Product Description 12", QuantityInStock: 10, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd1/h2f/64255724159006.jpg?format=preview-large"},
		{Name: "Apple iPhone 15 128Gb Blue", Price: 397652, Description: "Product Description 13", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd1/h07/83559339032606.png?format=preview-large"},
		{Name: "Xiaomi Redmi 12C 4GB/128GB Gray", Price: 48800, Description: "Product Description 14", QuantityInStock: 16, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/ha6/h82/69586957697054.jpg?format=preview-large"},
		{Name: "Poco X6 Pro 12GB/512GB Black", Price: 172598, Description: "Product Description 15", QuantityInStock: 7, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hbe/h45/84940376899614.jpg?format=preview-large"},
		{Name: "Apple iPhone 15 128Gb Black", Price: 401794, Description: "Product Description 3", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/he2/h1d/83559338442782.png?format=preview-large"},
	}
	err = product.InsertInitialProducts(db, products)
	if err != nil {
		log.Fatal(err)
	}

	err = user.EnsureUserTableExists(db)
	if err != nil {
		log.Fatal(err)
	}

	err = cart.EnsureCartTableExists(db)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
		handlers.LogoutHandler(w, r)
	}).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/products", handlers.GetAllProducts(db)).Methods("GET")
	r.HandleFunc("/api/products/{id}", handlers.GetProductByID(db)).Methods("GET")
	r.HandleFunc("/api/products", handlers.CreateProduct(db)).Methods("POST")
	r.HandleFunc("/api/products/{id}", handlers.UpdateProduct(db)).Methods("PUT")
	r.HandleFunc("/api/products/{id}", handlers.DeleteProduct(db)).Methods("DELETE")

	r.HandleFunc("/api/cart/{product_id}", handlers.GetCartItem(db)).Methods("GET")
	r.HandleFunc("/api/cart/{product_id}", handlers.AddToCart(db)).Methods("POST")
	r.HandleFunc("/api/cart/{product_id}", handlers.RemoveFromCart(db)).Methods("DELETE")

	// Protected routes

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(r)

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}

}
