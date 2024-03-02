// internal/handlers/product_handler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"database/sql"
	"onlinestore/pkg/product"
)

func GetProducts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	products, err := product.GetAllProductsFromDB(db)
	if err != nil {
		http.Error(w, "Error retrieving products", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
