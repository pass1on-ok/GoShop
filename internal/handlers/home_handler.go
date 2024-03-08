// internal/handlers/home_handler.go
package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"onlinestore/pkg/product"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	products, err := product.GetAllProductsFromDB(db)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error retrieving products:", err)
		return
	}

	tmpl, err := template.ParseFiles("web-page/homepage/home.html")
	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
		return
	}
}
