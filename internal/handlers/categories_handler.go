package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"onlinestore/pkg/category"
	"onlinestore/pkg/product"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllCategories(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := category.GetAllCategoriesFromDB(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
	}
}

func GetCategoryByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		categoryID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		c, err := category.GetCategoryByIDFromDB(db, categoryID)
		if err != nil {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}

		jsonData, err := json.Marshal(c)
		if err != nil {
			http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func CreateCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCategory category.Category
		err := json.NewDecoder(r.Body).Decode(&newCategory)
		if err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}

		err = category.InsertCategoryToDB(db, newCategory)
		if err != nil {
			http.Error(w, "Error adding category", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		categoryID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		var updatedCategory category.Category
		err = json.NewDecoder(r.Body).Decode(&updatedCategory)
		if err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}

		updatedCategory.ID = categoryID
		err = category.UpdateCategoryInDB(db, updatedCategory)
		if err != nil {
			http.Error(w, "Error updating category", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		categoryID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		err = category.DeleteCategoryFromDB(db, categoryID)
		if err != nil {
			http.Error(w, "Error deleting category", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetProductsByCategoryID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		categoryID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		products, err := product.GetProductsByCategoryIDFromDB(db, categoryID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
