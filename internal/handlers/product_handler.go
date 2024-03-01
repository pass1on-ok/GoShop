package handlers

import (
	"encoding/json"
	"net/http"

	"onlinestore/pkg/product"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := product.GetAllProducts()
	jsonData, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
