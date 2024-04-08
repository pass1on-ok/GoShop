// internal/handlers/home_handler.go
package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"onlinestore/pkg/product"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Retrieve page and perPage parameters from query string
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("perPage")

	// Convert page and perPage parameters to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // Default to page 1 if invalid or not provided
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 6 // Default to 6 products per page if invalid or not provided
	}

	// Retrieve products for the specified page
	products, err := product.GetProductsByPageFromDB(db, page, perPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error retrieving products: %v", err)
		return
	}

	// Render the HTML template with the retrieved products
	tmpl, err := template.ParseFiles("web-page/homepage/home.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error parsing HTML template: %v", err)
		return
	}

	// Execute the template with the products data
	err = tmpl.Execute(w, products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error executing HTML template: %v", err)
		return
	}
}
