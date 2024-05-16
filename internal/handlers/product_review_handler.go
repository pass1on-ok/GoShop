//internal\handlers\product_review_handler.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/review"

	"strconv"

	"github.com/gorilla/mux"
)

// CreateProductReview создает новый отзыв о продукте
func CreateProductReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newReview review.Review
		err := json.NewDecoder(r.Body).Decode(&newReview)
		if err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}

		// Получаем ID продукта из параметра маршрута
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		// Устанавливаем product_id в отзыве
		newReview.ProductID = productID

		// Получаем текущий user_id из контекста или сессии
		userID := getCurrentUserIDFromContextOrSession(r)
		if userID == 0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err = review.InsertProductReviewToDB(db, newReview, userID)
		if err != nil {
			http.Error(w, "Error adding product review", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// GetProductReviews возвращает отзывы о продукте по его ID
func GetProductReviews(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		reviews, err := review.GetProductReviewsFromDB(db, productID)
		if err != nil {
			http.Error(w, "Error retrieving product reviews", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
	}
}
