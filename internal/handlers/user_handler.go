// internal\handlers\user_handler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"onlinestore/pkg/user"
)

func ProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := getCurrentUserIDFromContextOrSession(r)

		userInfo, err := user.GetUserByIDFromDB(db, userID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to fetch user profile: %v", err), http.StatusInternalServerError)
			return
		}

		// Mask the password before sending the user info
		//userInfo.Password = ""

		// Convert user info to JSON and send response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(userInfo); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode user profile to JSON: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
