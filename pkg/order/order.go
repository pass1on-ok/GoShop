// pkg/order/order.go
package order

import (
	"database/sql"
	"log"
)

func EnsureOrderTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Println("Error creating orders table:", err)
		return err
	}
	log.Println("Orders table created successfully")
	return nil
}
