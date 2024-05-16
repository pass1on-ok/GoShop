// internal/handlers/order_handler.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/order"
	"strconv"

	"github.com/gorilla/mux"
)

func PostOrderHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	type RequestBody struct {
		ProductID int `json:"product_id"`
	}
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := getCurrentUserIDFromContextOrSession(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	orderID, err := order.CreateOrder(userID, requestBody.ProductID, db)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]int{"order_id": orderID}
	json.NewEncoder(w).Encode(response)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	params := mux.Vars(r)
	orderID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderInfo, err := order.GetOrderByID(orderID, db)
	if err != nil {
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderInfo)
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	orders, err := order.GetOrders(db)
	if err != nil {
		http.Error(w, "Failed to get orders", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"orders": orders}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
