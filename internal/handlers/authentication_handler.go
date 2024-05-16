// handlers/authentication_handler.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"onlinestore/pkg/auth"
	"onlinestore/pkg/user"
	"strings"

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

	user, err := auth.GetUserByUsername(credentials.Username, db)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	err = auth.UpdateUserToken(user.ID, token, db)
	if err != nil {
		http.Error(w, "Failed to update user token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"token": token}
	json.NewEncoder(w).Encode(response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Token is missing", http.StatusUnauthorized)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := auth.ParseToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = auth.UpdateUserToken(claims.UserID, "", db)
	if err != nil {
		http.Error(w, "Failed to delete token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User logged out successfully"}
	json.NewEncoder(w).Encode(response)
}

func getCurrentUserIDFromContextOrSession(r *http.Request) int {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		fmt.Println("Token is missing")
		return 0
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return 0 // Если произошла ошибка парсинга токена, вернем 0 или другое значение по умолчанию
	}
	if !token.Valid {
		fmt.Println("Token is invalid")
		return 0 // Если токен недействителен, вернем 0 или другое значение по умолчанию
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Failed to get claims from token")
		return 0 // Если не удалось получить claims из токена, вернем 0 или другое значение по умолчанию
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		fmt.Println("Failed to convert user_id to float64")
		return 0 // Если не удалось преобразовать user_id в число float64, вернем 0 или другое значение по умолчанию
	}

	return int(userID) // Возвращаем user_id как целочисленное значение
}
