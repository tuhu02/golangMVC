package main

import (
	"belajar/golangMVC/config"
	"belajar/golangMVC/controllers/categorycontroller"
	"belajar/golangMVC/controllers/homecontroller"
	"belajar/golangMVC/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Serve static files (CSS, JS, images)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css/"))))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("views/dist/"))))

	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/update", categorycontroller.Update)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Product routes - CRUD lengkap untuk manajemen produk
	http.HandleFunc("/products", productcontroller.Index)           // Menampilkan daftar produk
	http.HandleFunc("/products/detail", productcontroller.Detail)   // Menampilkan detail produk
	http.HandleFunc("/products/add", productcontroller.Add)         // Form tambah & proses create produk
	http.HandleFunc("/products/update", productcontroller.Update)   // Form edit & proses update produk
	http.HandleFunc("/products/delete", productcontroller.Delete)   // Menghapus produk

	log.Println("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
