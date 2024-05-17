// pkg/category/category.go
package category

import (
	"database/sql"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func InsertInitialCategories(db *sql.DB, categories []Category) error {
	for _, category := range categories {
		exists, err := CategoryExistsByName(db, category.Name)
		if err != nil {
			return err
		}

		if !exists {
			err := InsertCategoryToDB(db, category)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CategoryExistsByName проверяет, существует ли категория с данным именем.
func CategoryExistsByName(db *sql.DB, name string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM categories WHERE name = $1)", name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetAllCategoriesFromDB(db *sql.DB) ([]Category, error) {
	var categories []Category

	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func GetCategoryByIDFromDB(db *sql.DB, id int) (*Category, error) {
	var c Category

	row := db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id)
	err := row.Scan(&c.ID, &c.Name)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func InsertCategoryToDB(db *sql.DB, category Category) error {
	_, err := db.Exec("INSERT INTO categories (name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategoryInDB(db *sql.DB, c Category) error {
	_, err := db.Exec("UPDATE categories SET name = $1 WHERE id = $2", c.Name, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategoryFromDB(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func EnsureCategoryTableExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}
