// internal\handlers\filter_handler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/product"
)

func GetFilteredProductsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	filter := r.URL.Query().Get("filter")
	products, err := product.GetFilteredProducts(db, filter)
	if err != nil {
		http.Error(w, "Failed to fetch products: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
