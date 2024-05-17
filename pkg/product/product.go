// pkg/product/product.go
package product

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ID              int
	Name            string
	Price           float64
	Description     string
	QuantityInStock int
	ImagePath       string
	CategoryID      int
}

func GetAllProductsFromDB(db *sql.DB) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, name, price, description, quantity_in_stock, imagepath, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func GetProductByIDFromDB(db *sql.DB, id int) (*Product, error) {
	var p Product

	row := db.QueryRow("SELECT id, name, price, description, quantity_in_stock, imagepath, category_id FROM products WHERE id = $1", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath, &p.CategoryID)
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
	_, err := db.Exec(`INSERT INTO products (name, price, description, quantity_in_stock, imagepath, category_id)
						VALUES ($1, $2, $3, $4, $5, $6)`,
		product.Name, product.Price, product.Description, product.QuantityInStock, product.ImagePath, product.CategoryID)
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
	_, err := db.Exec("UPDATE products SET name = $1, price = $2, description = $3, quantity_in_stock = $4, imagepath = $5, category_id = $6 WHERE id = $7",
		p.Name, p.Price, p.Description, p.QuantityInStock, p.ImagePath, p.CategoryID, p.ID)
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
		imagepath TEXT,
		
		category_id INTEGER REFERENCES categories(id)
	)`)
	return err
}

func GetProductsByCategoryIDFromDB(db *sql.DB, categoryID int) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, name, price, description, quantity_in_stock, imagepath, category_id FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func GetPaginatedProducts(db *sql.DB, pageSize, offset int) ([]Product, error) {
	var products []Product
	query := "SELECT id, name, price, description, quantity_in_stock, imagepath, category_id FROM products LIMIT $1 OFFSET $2"
	rows, err := db.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func GetSortedProducts(db *sql.DB, sortBy, sortOrder string) ([]Product, error) {
	var products []Product
	query := fmt.Sprintf("SELECT id, name, price, description, quantity_in_stock, imagepath, category_id FROM products ORDER BY %s %s", sortBy, sortOrder)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func GetFilteredProducts(db *sql.DB, filter string) ([]Product, error) {
	var products []Product
	query := "SELECT id, name, price, description, quantity_in_stock, imagepath, category_id FROM products WHERE name ILIKE '%' || $1 || '%' OR description ILIKE '%' || $1 || '%'"
	rows, err := db.Query(query, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.QuantityInStock, &p.ImagePath, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}
