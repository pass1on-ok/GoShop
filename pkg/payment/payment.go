// pkg/payment/payment.go

package payment

import (
	"database/sql"
	"time"
)

type PaymentInfo struct {
	ID            int       `json:"id"`
	OrderID       int       `json:"order_id"`
	PaymentAmount float64   `json:"payment_amount"`
	PaymentDate   time.Time `json:"payment_date"`
	UserID        int       `json:"user_id"`
}

func EnsurePaymentInfoTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS payment_info (
        id SERIAL PRIMARY KEY,
        order_id INT NOT NULL,
		user_id INT NOt NULL,
        payment_amount NUMERIC NOT NULL,
        payment_date TIMESTAMP NOT NULL,
        FOREIGN KEY (order_id) REFERENCES orders(id),
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (payment_amount) REFERENCES products(id)
    )`)
	if err != nil {
		return err
	}
	return nil
}

func CreatePaymentInfo(db *sql.DB, payment PaymentInfo) error {
	_, err := db.Exec(`INSERT INTO payment_info (order_id, payment_amount, payment_date, user_id) VALUES ($1, $2, $3, $4)`,
		payment.OrderID, payment.PaymentAmount, payment.PaymentDate, payment.UserID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPaymentsByUserID(db *sql.DB, userID int) ([]PaymentInfo, error) {

	var payments []PaymentInfo

	rows, err := db.Query("SELECT id, order_id, payment_amount, payment_date FROM payment_info WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment PaymentInfo
		err := rows.Scan(&payment.ID, &payment.OrderID, &payment.PaymentAmount, &payment.PaymentDate)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

func GetPaymentInfoByID(db *sql.DB, ID, userID int) (*PaymentInfo, error) {

	var payment PaymentInfo

	err := db.QueryRow("SELECT id, order_id, payment_amount, payment_date FROM payment_info WHERE id = $1 AND user_id = $2", ID, userID).
		Scan(&payment.ID, &payment.OrderID, &payment.PaymentAmount, &payment.PaymentDate)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
