// pkg/payment/payment.go
package payment

import (
	"database/sql"
	"time"
)

type PaymentInfo struct {
	PaymentID     int       `json:"payment_id"`
	OrderID       int       `json:"order_id"`
	PaymentAmount float64   `json:"payment_amount"`
	PaymentDate   time.Time `json:"payment_date"`
	UserID        int       `json:"user_id"`
}

func EnsurePaymentInfoTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS payment_info (
        payment_id SERIAL PRIMARY KEY,
        order_id INT NOT NULL,
		user_id INT NOt NULL,
        payment_amount NUMERIC NOT NULL,
        payment_date TIMESTAMP NOT NULL,
        FOREIGN KEY (order_id) REFERENCES orders(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
    )`)
	if err != nil {
		return err
	}
	return nil
}

func CreatePaymentInfo(db *sql.DB, payment PaymentInfo) error {
	_, err := db.Exec(`INSERT INTO payment_info (order_id, payment_amount, payment_date) VALUES ($1, $2, $3)`,
		payment.OrderID, payment.PaymentAmount, payment.PaymentDate)
	if err != nil {
		return err
	}
	return nil
}

func GetPaymentInfoByID(db *sql.DB, paymentID int) (*PaymentInfo, error) {
	var payment PaymentInfo

	err := db.QueryRow(`SELECT payment_id, order_id, payment_amount, payment_date FROM payment_info WHERE payment_id = $1`, paymentID).Scan(
		&payment.PaymentID, &payment.OrderID, &payment.PaymentAmount, &payment.PaymentDate)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
