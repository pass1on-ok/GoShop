package product

import (
	"database/sql"
)

type Product struct {
	ID              int
	Name            string
	Price           float64
	Description     string
	QuantityInStock int
}

func GetAllProductsFromDB(db *sql.DB) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, name, price, description, quantity_in_stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
