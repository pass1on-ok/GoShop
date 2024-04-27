// internal/handlers/product_handler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/product"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := product.GetAllProductsFromDB(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func GetProductByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product, err := product.GetProductByIDFromDB(db, productID)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		jsonData, err := json.Marshal(product)
		if err != nil {
			http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newProduct product.Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}

		err = product.InsertProductToDB(db, newProduct)
		if err != nil {
			http.Error(w, "Error adding product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var updatedDescription struct {
			Description string `json:"description"`
		}
		err = json.NewDecoder(r.Body).Decode(&updatedDescription)
		if err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}

		currentProduct, err := product.GetProductByIDFromDB(db, productID)
		if err != nil {
			http.Error(w, "Error retrieving product", http.StatusInternalServerError)
			return
		}

		currentProduct.Description = updatedDescription.Description

		err = product.UpdateProductInDB(db, *currentProduct)
		if err != nil {
			http.Error(w, "Error updating product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		err = product.DeleteProductFromDB(db, productID)
		if err != nil {
			http.Error(w, "Error deleting product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
