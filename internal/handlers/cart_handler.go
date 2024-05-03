// handlers.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/cart"
)

type CartHandler struct {
	DB *sql.DB
}

type CartItemRequest struct {
	UserID    int `json:"userID"`
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

func NewCartHandler(db *sql.DB) *CartHandler {
	return &CartHandler{
		DB: db,
	}
}

func (h *CartHandler) AddItemToCartHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody CartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	item := cart.Item{
		ProductID: requestBody.ProductID,
		Quantity:  requestBody.Quantity,
	}

	err := cart.AddToCart(h.DB, item)
	if err != nil {
		http.Error(w, "Failed to add item to cart: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Item added to cart"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
