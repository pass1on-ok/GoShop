//internal\handlers\sort_handler.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/product"
)

func GetSortedProductsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = "price"
	}
	sortOrder := r.URL.Query().Get("sortOrder")
	if sortOrder == "" {
		sortOrder = "ASC"
	}

	products, err := product.GetSortedProducts(db, sortBy, sortOrder)
	if err != nil {
		http.Error(w, "Failed to fetch products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
