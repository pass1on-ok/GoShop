package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yourapp/db"
	"yourapp/models"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user input (e.g., check for required fields)

	// Hash user password before saving it to the database

	// Save user to the database
	err = db.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func main() {
	// Initialize database connection
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// Define HTTP routes
	http.HandleFunc("/register", RegisterUserHandler)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
