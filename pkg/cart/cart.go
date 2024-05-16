// pkg/cart/cart.go
package cart

import (
	"database/sql"
)

type CartItem struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func EnsureCartTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS cart (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        product_id INT NOT NULL,
        quantity INT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (product_id) REFERENCES products(id)
    )`)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemByProductID(db *sql.DB, userID, productID int) (*CartItem, error) {
	var item CartItem
	err := db.QueryRow("SELECT id, user_id, product_id, quantity FROM cart WHERE user_id = $1 AND product_id = $2", userID, productID).Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func AddProductToCart(db *sql.DB, userID, productID, quantity int) error {
	_, err := db.Exec("INSERT INTO cart (user_id, product_id, quantity) VALUES ($1, $2, $3)", userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func RemoveProductFromCart(db *sql.DB, userID, productID int) error {
	_, err := db.Exec("DELETE FROM cart WHERE user_id = $1 AND product_id = $2", userID, productID)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemsByUserID(db *sql.DB, userID int) ([]CartItem, error) {
	rows, err := db.Query("SELECT id, user_id, product_id, quantity FROM cart WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CartItem
	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
