// cmd/main.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"onlinestore/internal/handlers"
	"onlinestore/pkg/cart"
	"onlinestore/pkg/category"
	"onlinestore/pkg/order"
	"onlinestore/pkg/payment"
	"onlinestore/pkg/product"
	"onlinestore/pkg/review"
	"onlinestore/pkg/user"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:kumar@localhost/Online%20Store?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = category.EnsureCategoryTableExists(db)
	if err != nil {
		log.Fatal(err)
	}

	categories := []category.Category{
		{Name: "Home Appliances"},
		{Name: "Phones"},
		{Name: "End Devices"},
	}

	err = category.InsertInitialCategories(db, categories)
	if err != nil {
		log.Fatal(err)
	}

	err = product.EnsureTableExists(db)
	if err != nil {
		log.Fatal(err)
	}

	products := []product.Product{
		{Name: "Apple iPhone 15 Pro Max 256Gb Gray", Price: 654857, Description: "Product Description 1", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hc1/h65/83559848181790.png?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 13 128Gb Midnight Black", Price: 299929, Description: "Product Description 2", QuantityInStock: 3, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h32/h70/84378448199710.jpg?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 15 128Gb Black", Price: 401794, Description: "Product Description 3", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/he2/h1d/83559338442782.png?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 11 128Gb Slim Box Black", Price: 244900, Description: "Product Description 4", QuantityInStock: 2, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd8/hac/63897052413982.jpg?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 13 128Gb White", Price: 300000, Description: "Product Description 5", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hc9/h90/64209083007006.jpg?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 14 128Gb Black", Price: 341178, Description: "Product Description 6", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h98/h2b/64400497737758.jpg?format=preview-large", CategoryID: 2},
		{Name: "Samsung Galaxy A24 6GB/128GB Black", Price: 91636, Description: "Product Description 7", QuantityInStock: 20, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hdc/h12/80750151303198.jpg?format=preview-large", CategoryID: 2},
		{Name: "Xiaomi Redmi 12 4G 8GB/256GB Black", Price: 68759, Description: "Product Description 8", QuantityInStock: 10, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h75/hbc/81335343775774.png?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 15 Pro 256Gb Blue", Price: 597788, Description: "Product Description 9", QuantityInStock: 30, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h88/hde/83559835697182.jpg?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 14 128Gb Starlight", Price: 341100, Description: "Product Description 10", QuantityInStock: 2, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h9f/h49/64481569832990.jpg?format=preview-large", CategoryID: 2},
		{Name: "Samsung Galaxy A54 5G 8GB/256GB Black", Price: 186756, Description: "Product Description 11", QuantityInStock: 11, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/h81/h13/80435536887838.jpg?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 13 128Gb Green", Price: 307856, Description: "Product Description 12", QuantityInStock: 10, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd1/h2f/64255724159006.jpg?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 15 128Gb Blue", Price: 397652, Description: "Product Description 13", QuantityInStock: 12, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hd1/h07/83559339032606.png?format=preview-large", CategoryID: 2},
		{Name: "Xiaomi Redmi 12C 4GB/128GB Gray", Price: 48800, Description: "Product Description 14", QuantityInStock: 16, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/ha6/h82/69586957697054.jpg?format=preview-large", CategoryID: 2},
		{Name: "Poco X6 Pro 12GB/512GB Black", Price: 172598, Description: "Product Description 15", QuantityInStock: 7, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/hbe/h45/84940376899614.jpg?format=preview-large", CategoryID: 2},
		{Name: "Apple iPhone 15 128Gb Black", Price: 401794, Description: "Product Description 3", QuantityInStock: 1, ImagePath: "https://resources.cdn-kaspi.kz/img/m/p/he2/h1d/83559338442782.png?format=preview-large", CategoryID: 2},
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

	err = order.EnsureOrderTableExists(db)
	if err != nil {
		log.Fatal(err)
	}

	err = payment.EnsurePaymentInfoTableExists(db)
	if err != nil {
		log.Fatal(err)
	}

	err = review.EnsureReviewTableExists(db)
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
		handlers.LogoutHandler(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/products", handlers.GetAllProducts(db)).Methods("GET")
	r.HandleFunc("/api/products/{id}", handlers.GetProductByID(db)).Methods("GET")
	r.HandleFunc("/api/products", handlers.CreateProduct(db)).Methods("POST")
	r.HandleFunc("/api/products/{id}", handlers.UpdateProduct(db)).Methods("PUT")
	r.HandleFunc("/api/products/{id}", handlers.DeleteProduct(db)).Methods("DELETE")

	r.HandleFunc("/api/products/{id}/review", handlers.CreateProductReview(db)).Methods("POST")
	r.HandleFunc("/api/products/{id}/review", handlers.GetProductReviews(db)).Methods("GET")

	r.HandleFunc("/api/products/{id}/cart", handlers.AddToCartForProduct(db)).Methods("POST")

	r.HandleFunc("/api/cart", handlers.GetCartItemsHandler(db)).Methods("GET")
	r.HandleFunc("/api/cart/{product_id}", handlers.AddToCart(db)).Methods("POST")
	r.HandleFunc("/api/cart/{product_id}", handlers.RemoveFromCart(db)).Methods("DELETE")

	r.HandleFunc("/api/payments", handlers.GetAllPaymentsForCurrentUser(db)).Methods("GET")
	r.HandleFunc("/api/payments/{id}", handlers.GetPaymentByID(db)).Methods("GET")
	r.HandleFunc("/api/payments", handlers.CreatePayment(db)).Methods("POST")

	r.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostOrderHandler(w, r, db)
	}).Methods("POST")
	r.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetOrdersHandler(w, r, db)
	}).Methods("GET")
	r.HandleFunc("/api/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetOrderHandler(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/orders/{id}/payment", handlers.CreatePaymentForOrder(db)).Methods("POST")

	r.HandleFunc("/api/profile", handlers.ProfileHandler(db)).Methods("GET")

	r.HandleFunc("/api/categories", handlers.GetAllCategories(db)).Methods("GET")
	r.HandleFunc("/api/categories/{id}", handlers.GetCategoryByID(db)).Methods("GET")
	r.HandleFunc("/api/categories", handlers.CreateCategory(db)).Methods("POST")
	r.HandleFunc("/api/categories/{id}", handlers.UpdateCategory(db)).Methods("PUT")
	r.HandleFunc("/api/categories/{id}", handlers.DeleteCategory(db)).Methods("DELETE")

	r.HandleFunc("/api/categories/{id}/products", handlers.GetProductsByCategoryID(db)).Methods("GET")

	r.HandleFunc("/api/products/pagination/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetPaginatedProductsHandler(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/products/sort/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSortedProductsHandler(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/products/filter/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetFilteredProductsHandler(w, r, db)
	}).Methods("GET")

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
