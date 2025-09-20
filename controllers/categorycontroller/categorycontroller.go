package categorycontroller

import (
	"belajar/golangMVC/entities"
	"belajar/golangMVC/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		// Validasi input
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		var category entities.Category
		category.Name = name
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			http.Error(w, "Failed to create category", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
		return
	}
}


func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = categorymodel.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/update.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}

		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		category, err := categorymodel.Detail(id)
		if err != nil {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}

		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
		return
	}

	if r.Method == "POST" {
		// Validasi ID
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Validasi name
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		var category entities.Category
		category.Id = uint(id)
		category.Name = name
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Update(id, category); !ok {
			http.Error(w, "Failed to update category", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
