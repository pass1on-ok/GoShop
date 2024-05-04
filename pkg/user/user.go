// pkg/user/users.go

package user

import (
	"database/sql"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Token    string
}

func CreateUser(db *sql.DB, newUser User) error {
	_, err := db.Exec(`INSERT INTO users (username, email, password, token) VALUES ($1, $2, $3, $4)`,
		newUser.Username, newUser.Email, newUser.Password, newUser.Token)
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
        password TEXT NOT NULL,
        token TEXT
    )`)
	if err != nil {
		return err
	}
	return nil
}
