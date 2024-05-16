// handlers/payment_handler.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/order"
	"onlinestore/pkg/payment"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type CreatePaymentRequest struct {
	OrderID       int     `json:"order_id"`
	PaymentAmount float64 `json:"payment_amount"`
}

func GetPayment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		paymentID, err := strconv.Atoi(params["payment_id"])
		if err != nil {
			http.Error(w, "Invalid payment ID", http.StatusBadRequest)
			return
		}

		p, err := payment.GetPaymentInfoByID(db, paymentID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Payment not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error retrieving payment", http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(p)
	}
}

func CreatePayment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody CreatePaymentRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		paymentInfo := payment.PaymentInfo{
			OrderID:       requestBody.OrderID,
			PaymentAmount: requestBody.PaymentAmount,
			PaymentDate:   time.Now(),
		}

		err = payment.CreatePaymentInfo(db, paymentInfo)
		if err != nil {
			http.Error(w, "Error creating payment", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func CreatePaymentForOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orderID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		// Получение user_id из сеанса или токена
		userID := getCurrentUserIDFromContextOrSession(r)

		// Получение цены продуктов для данного заказа из базы данных
		orderTotal, err := order.GetOrderTotal(db, orderID)
		if err != nil {
			http.Error(w, "Error retrieving order total", http.StatusInternalServerError)
			return
		}

		// Создание информации о платеже
		paymentInfo := payment.PaymentInfo{
			OrderID:       orderID,
			PaymentAmount: orderTotal,
			PaymentDate:   time.Now(),
			UserID:        userID, // Передача user_id в информацию о платеже
		}

		// Сохранение информации о платеже в базе данных
		err = payment.CreatePaymentInfo(db, paymentInfo)
		if err != nil {
			http.Error(w, "Error creating payment", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
