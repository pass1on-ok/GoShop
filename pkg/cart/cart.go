// cart.go
package cart

import (
	"database/sql"
	"errors"
	"log"
)

type Item struct {
	ProductID int
	Quantity  int
}

func AddToCart(db *sql.DB, item Item) error {

	var quantityInStock int
	err := db.QueryRow("SELECT quantity_in_stock FROM products WHERE id = $1", item.ProductID).Scan(&quantityInStock)
	if err != nil {
		return err
	}

	if quantityInStock < item.Quantity {
		return errors.New("insufficient quantity in stock")
	}

	_, err = db.Exec("INSERT INTO cart (product_id, quantity) VALUES ($1, $2)", item.ProductID, item.Quantity)
	if err != nil {
		return err
	}

	log.Printf("Added item to cart: %+v\n", item)
	return nil
}
