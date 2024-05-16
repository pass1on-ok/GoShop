// pkg/order/order.go

package order

import (
	"database/sql"
	"log"
)

type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	OrderDate string `json:"order_date"`
}

func EnsureOrderTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		product_id INT NOT NULL,
		order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (product_id) REFERENCES products(id)
	)`)
	if err != nil {
		log.Println("Error creating orders table:", err)
		return err
	}
	log.Println("Orders table created successfully")
	return nil
}

func CreateOrder(userID, productID int, db *sql.DB) (int, error) {
	var orderID int
	err := db.QueryRow("INSERT INTO orders (user_id, product_id) VALUES ($1, $2) RETURNING id", userID, productID).Scan(&orderID)
	if err != nil {
		log.Println("Error creating order:", err)
		return 0, err
	}
	return orderID, nil
}

func GetOrders(db *sql.DB) ([]Order, error) {
	rows, err := db.Query("SELECT id, user_id, order_date FROM orders")
	if err != nil {
		log.Println("Error querying orders:", err)
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate)
		if err != nil {
			log.Println("Error scanning order:", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over orders:", err)
		return nil, err
	}
	return orders, nil
}

// GetOrderByID извлекает заказ из базы данных по его ID
func GetOrderByID(orderID int, db *sql.DB) (*Order, error) {
	var order Order
	err := db.QueryRow("SELECT id, user_id, order_date FROM orders WHERE id = $1", orderID).Scan(&order.ID, &order.UserID, &order.OrderDate)
	if err != nil {
		log.Println("Error querying order by ID:", err)
		return nil, err
	}
	return &order, nil
}

func GetOrderTotal(db *sql.DB, orderID int) (float64, error) {
	// Выполнение запроса к базе данных для получения цен продуктов в заказе
	rows, err := db.Query("SELECT SUM(price) FROM products WHERE id IN (SELECT product_id FROM order_items WHERE order_id = $1)", orderID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var orderTotal float64
	for rows.Next() {
		err := rows.Scan(&orderTotal)
		if err != nil {
			return 0, err
		}
	}

	return orderTotal, nil
}
