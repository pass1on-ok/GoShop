// pkg/product/product.go
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

	rows, err := db.Query("SELECT id, name, price, description, quantity_in_stock, imagepath FROM products")
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

func GetProductByIDFromDB(db *sql.DB, id int) (*Product, error) {
	var p Product

	row := db.QueryRow("SELECT id, name, price, description, quantity_in_stock, imagepath FROM products WHERE id = $1", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func InsertInitialProducts(db *sql.DB, products []Product) error {
	for _, product := range products {
		exists, err := ProductExistsByParams(db, product)
		if err != nil {
			return err
		}

		if !exists {
			err := InsertProductToDB(db, product)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func InsertProductToDB(db *sql.DB, product Product) error {
	_, err := db.Exec(`INSERT INTO products (name, price, description, quantity_in_stock, imagepath)
						VALUES ($1, $2, $3, $4, $5)`,
		product.Name, product.Price, product.Description, product.QuantityInStock, product.ImagePath)
	if err != nil {
		return err
	}
	return nil
}

func ProductExistsByParams(db *sql.DB, product Product) (bool, error) {
	var exists bool
	err := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM products 
		WHERE name = $1 AND price = $2 AND description = $3 
		AND quantity_in_stock = $4 AND imagepath = $5)`,
		product.Name, product.Price, product.Description, product.QuantityInStock, product.ImagePath).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func UpdateProductInDB(db *sql.DB, p Product) error {
	_, err := db.Exec("UPDATE products SET name = $1, price = $2, description = $3, quantity_in_stock = $4, imagepath = $5 WHERE id = $6",
		p.Name, p.Price, p.Description, p.QuantityInStock, p.ImagePath, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductFromDB(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func EnsureTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT,
		price NUMERIC,
		description TEXT,
		quantity_in_stock INTEGER,
		imagepath TEXT
	)`)
	return err
}
