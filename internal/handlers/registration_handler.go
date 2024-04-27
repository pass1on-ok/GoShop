// internal\handlers\registration_handler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/user"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
	/*
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User created successfully")
	*/
}
