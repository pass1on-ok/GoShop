//pkg/product/product.go

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
	ImagePath       string
}

func GetAllProductsFromDB(db *sql.DB) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, name, price, description, quantity_in_stock,imagepath FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func AddProductsToDB(db *sql.DB, products []Product) error {
	for _, product := range products {

		exists, err := ProductExists(db, product.Name)
		if err != nil {
			return err
		}

		if !exists {
			err := AddProductToDB(db, product)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func AddProductToDB(db *sql.DB, product Product) error {
	_, err := db.Exec(`INSERT INTO products (name, price, description, quantity_in_stock, imagepath)
						VALUES ($1, $2, $3, $4, $5)`,
		product.Name, product.Price, product.Description, product.QuantityInStock, product.ImagePath)
	if err != nil {
		return err
	}
	return nil
}

func ProductExists(db *sql.DB, productName string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM products WHERE name = $1)", productName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
