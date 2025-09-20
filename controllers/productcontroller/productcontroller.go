package productcontroller

import (
	"belajar/golangMVC/entities"
	"belajar/golangMVC/models/categorymodel"
	"belajar/golangMVC/models/productmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// Index menampilkan daftar semua produk
func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/Product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

// Detail menampilkan detail produk berdasarkan ID
func Detail(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product, err := productmodel.Detail(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/Product/detail.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}

// Add menampilkan form tambah produk (GET) dan memproses penambahan (POST)
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Ambil semua kategori untuk dropdown
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp, err := template.ParseFiles("views/Product/create.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		return
	}

	if r.Method == "POST" {
		// Validasi input
		name := r.FormValue("name")
		categoryIdStr := r.FormValue("category_id")
		stockStr := r.FormValue("stock")
		description := r.FormValue("description")

		if name == "" || categoryIdStr == "" || stockStr == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		categoryId, err := strconv.Atoi(categoryIdStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		stock, err := strconv.ParseInt(stockStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid stock value", http.StatusBadRequest)
			return
		}

		var product entities.Product
		product.Name = name
		product.CategoryId = uint(categoryId)
		product.Stock = stock
		product.Description = description
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			http.Error(w, "Failed to create product", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}
}

// Update menampilkan form edit produk (GET) dan memproses update (POST)
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID parameter is required", http.StatusBadRequest)
			return
		}
		
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		product, err := productmodel.Detail(id)
		if err != nil {
			http.Error(w, "Product not found: "+err.Error(), http.StatusNotFound)
			return
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"product":    product,
			"categories": categories,
		}

		temp, err := template.ParseFiles("views/Product/update.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = temp.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
			return
		}
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

		// Validasi input
		name := r.FormValue("name")
		categoryIdStr := r.FormValue("category_id")
		stockStr := r.FormValue("stock")
		description := r.FormValue("description")

		if name == "" || categoryIdStr == "" || stockStr == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		categoryId, err := strconv.Atoi(categoryIdStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		stock, err := strconv.ParseInt(stockStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid stock value", http.StatusBadRequest)
			return
		}

		var product entities.Product
		product.Id = uint(id)
		product.Name = name
		product.CategoryId = uint(categoryId)
		product.Stock = stock
		product.Description = description
		product.UpdatedAt = time.Now()

		if ok := productmodel.Update(id, product); !ok {
			http.Error(w, "Failed to update product", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

// Delete menghapus produk berdasarkan ID
func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = productmodel.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}