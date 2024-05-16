// handlers/cart_handler.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/cart"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCartItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["product_id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		userID := getCurrentUserIDFromContextOrSession(r)
		item, err := cart.GetCartItemByProductID(db, userID, productID)
		if err != nil {
			http.Error(w, "Error retrieving cart item", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(item)
	}
}

func GetCartItemsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем текущий user_id из контекста или сессии
		userID := getCurrentUserIDFromContextOrSession(r)
		if userID == 0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Получаем содержимое корзины для текущего пользователя
		items, err := cart.GetCartItemsByUserID(db, userID)
		if err != nil {
			http.Error(w, "Error retrieving cart items", http.StatusInternalServerError)
			return
		}

		// Отправляем содержимое корзины в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

type AddToCartRequest struct {
	UserID   int `json:"user_id"`
	Quantity int `json:"quantity"`
}

func AddToCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["product_id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		userID := getCurrentUserIDFromContextOrSession(r)

		var requestBody AddToCartRequest
		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		err = cart.AddProductToCart(db, userID, productID, requestBody.Quantity)
		if err != nil {
			http.Error(w, "Error adding product to cart", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveFromCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["product_id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		userID := getCurrentUserIDFromContextOrSession(r)
		err = cart.RemoveProductFromCart(db, userID, productID)
		if err != nil {
			http.Error(w, "Error removing product from cart", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func AddToCartForProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"]) // Получаем product_id из пути
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		// Получаем текущий user_id из контекста или сессии
		userID := getCurrentUserIDFromContextOrSession(r)

		var requestBody AddToCartRequest
		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Используем user_id и product_id для добавления товара в корзину
		err = cart.AddProductToCart(db, userID, productID, requestBody.Quantity)
		if err != nil {
			http.Error(w, "Error adding product to cart", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
