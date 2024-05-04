// handlers/authentication_handler.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/auth"
	"onlinestore/pkg/user"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newUser user.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid data format", http.StatusBadRequest)
		return
	}

	err = user.CreateUser(db, newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var credentials auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid data format", http.StatusBadRequest)
		return
	}

	token, err := auth.Login(credentials, db)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Получить пользователя из базы данных, чтобы обновить запись с токеном
	user, err := auth.GetUserByUsername(credentials.Username, db)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	// Обновить запись пользователя с токеном/

	err = auth.UpdateUserToken(user.ID, token, db)
	if err != nil {
		http.Error(w, "Failed to update user token", http.StatusInternalServerError)
		return
	}
	/*
		err = auth.UpdateUserToken(user.ID, token, db)
		if err != nil {
			http.Error(w, "Failed to update user token: "+err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"token": token}
	json.NewEncoder(w).Encode(response)
}

// LogoutHandler обрабатывает запрос на выход пользователя
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User logged out successfully"}
	json.NewEncoder(w).Encode(response)
}

func getCurrentUserIDFromContextOrSession(r *http.Request) int {
	// Получаем токен из заголовка запроса
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return 0 // Если токен не предоставлен, вернем 0
	}

	// Проверяем и парсим токен

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("JWT_SECRET")), nil // Здесь вместо "secret" должен быть ваш секретный ключ
	})
	if err != nil || !token.Valid {
		return 4 // Если токен недействителен, вернем 0
	}

	// Извлекаем user_id из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0 // Если не удалось получить claims из токена, вернем 0
	}

	userID, err := strconv.Atoi(claims["user_id"].(string))
	if err != nil {
		return 0 // Если не удалось преобразовать user_id в число, вернем 0
	}

	return userID

}
