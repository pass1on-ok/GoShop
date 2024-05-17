// pkg/user/users.go

package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
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

func GetUserByIDFromDB(db *sql.DB, userID int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email,password, token FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Token)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email, password, token FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Token)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUserToken(db *sql.DB, token string) error {
	_, err := db.Exec("UPDATE users SET token = $1 WHERE id = $2", token, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
