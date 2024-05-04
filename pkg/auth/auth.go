// pkg/auth/auth.go

package auth

import (
	"database/sql"
	"errors"
	"onlinestore/pkg/user"

	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret") // Секретный ключ для подписи токена

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GenerateToken создает новый JWT токен для пользователя
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Токен действителен 24 часа

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseToken парсит JWT токен и возвращает имя пользователя из него
func ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Username, nil
}

func Login(credentials Credentials, db *sql.DB) (string, error) {
	// Проверка учетных данных пользователя в базе данных
	user, err := GetUserByUsername(credentials.Username, db) // Передаем db здесь
	if err != nil {
		return "", err
	}

	// Проверка соответствия пароля
	if user.Password != credentials.Password {
		return "", errors.New("incorrect username or password")
	}

	// Генерация токена
	token, err := GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Дополнительные функции для работы с базой данных

func GetUserByUsername(username string, db *sql.DB) (*user.User, error) {
	// Выполнение запроса к базе данных для получения информации о пользователе по его имени
	var u user.User
	err := db.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		// Если пользователь не найден, возвращаем ошибку
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		// В случае другой ошибки возвращаем её
		return nil, err
	}
	// Возвращаем найденного пользователя и nil в качестве ошибки
	return &u, nil
}

func UpdateUserToken(userID int, token string, db *sql.DB) error {

	_, err := db.Exec("UPDATE users SET token = $1 WHERE id = $2", token, userID)
	if err != nil {
		return err
	}
	return nil
}

/*
func UpdateUserToken(userID int, token string, db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET token = $1 WHERE id = $2", token, userID)
	if err != nil {
		return err
	}
	return nil
}
*/
