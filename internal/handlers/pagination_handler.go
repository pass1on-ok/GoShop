// internal\handlers\pagination_handler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/product"
	"strconv"
)

func GetPaginatedProductsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 5
	}
	offset := (page - 1) * pageSize
	products, err := product.GetPaginatedProducts(db, pageSize, offset)
	if err != nil {
		http.Error(w, "Failed to fetch products: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
