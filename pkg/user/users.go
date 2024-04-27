package user

import (
	"database/sql"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	// Add more fields as needed
}

func CreateUser(db *sql.DB, newUser User) error {
	_, err := db.Exec(`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`,
		newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		return err
	}
	return nil
}

func EnsureUserTableExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`)
	if err != nil {
		return err
	}
	return nil
}
