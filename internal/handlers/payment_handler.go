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

func CreatePayment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody CreatePaymentRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		userID := getCurrentUserIDFromContextOrSession(r)

		paymentInfo := payment.PaymentInfo{
			OrderID:       requestBody.OrderID,
			PaymentAmount: requestBody.PaymentAmount,
			PaymentDate:   time.Now(),
			UserID:        userID,
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

		userID := getCurrentUserIDFromContextOrSession(r) // Получение user_id из текущего сеанса

		// Получение общей стоимости заказа
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
			UserID:        userID, // Использование user_id текущего сеанса
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

func GetAllPaymentsForCurrentUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := getCurrentUserIDFromContextOrSession(r) // Получение текущего пользователя из сессии

		payments, err := payment.GetAllPaymentsByUserID(db, userID)
		if err != nil {
			http.Error(w, "Error retrieving payments", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(payments)
	}
}

// Обработчик для получения платежа по его ID
func GetPaymentByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid payment ID", http.StatusBadRequest)
			return
		}
		userID := getCurrentUserIDFromContextOrSession(r)

		p, err := payment.GetPaymentInfoByID(db, ID, userID)
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
