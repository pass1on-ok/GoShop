// pkg/review/review.go
package review

import "database/sql"

type Review struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    int    `json:"user_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}

func EnsureReviewTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS product_reviews (
		id SERIAL PRIMARY KEY,
		product_id INT NOT NULL,
		user_id INT NOT NULL,
		rating INT NOT NULL,
		comment TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (product_id) REFERENCES products(id)

	)`)
	return err
}

func InsertProductReviewToDB(db *sql.DB, review Review, userID int) error {
	_, err := db.Exec("INSERT INTO product_reviews (product_id, user_id, rating, comment) VALUES ($1, $2, $3, $4)",
		review.ProductID, userID, review.Rating, review.Comment)
	if err != nil {
		return err
	}
	return nil
}

func GetProductReviewsFromDB(db *sql.DB, productID int) ([]Review, error) {
	rows, err := db.Query("SELECT id, product_id, user_id, rating, comment FROM product_reviews WHERE product_id = $1", productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(&review.ID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}
