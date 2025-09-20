package productmodel

import (
	"belajar/golangMVC/config"   // package untuk koneksi DB
	"belajar/golangMVC/entities" // package berisi struct Product, Category, dll
	"database/sql"
)

// GetAll mengambil semua data produk dari database
func GetAll() []entities.Product {
	// Query ke database: ambil data product dan nama kategori
	rows, err := config.DB.Query(`
		SELECT 
			products.id,
			products.name,
			products.category_id,
			categories.name AS category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
	`)
	if err != nil {
		panic(err) // hentikan program kalau query gagal
	}
	defer rows.Close() // pastikan koneksi rows ditutup setelah selesai

	// Siapkan slice untuk menampung hasil query
	var products []entities.Product

	// Looping tiap baris hasil query
	for rows.Next() {
		var product entities.Product

		// Masukkan data dari hasil query ke struct Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.CategoryId,
			&product.Category.Name, // akses nama kategori di struct Category
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			panic(err) // hentikan program kalau gagal scan
		}

		// Tambahkan product ke slice products
		products = append(products, product)
	}

	// Kembalikan slice berisi semua produk
	return products
}

// Create menambahkan produk baru ke database
func Create(product entities.Product) bool {
	result, err := config.DB.Exec(
		"INSERT INTO products(name, category_id, stock, description, created_at, updated_at) VALUES (?,?,?,?,?,?)",
		product.Name, product.CategoryId, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt,
	)
	if err != nil {
		return false
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return false
	}

	return lastInsertId > 0
}

// Update mengubah data produk berdasarkan ID
func Update(id int, product entities.Product) bool {
	result, err := config.DB.Exec(
		"UPDATE products SET name = ?, category_id = ?, stock = ?, description = ?, updated_at = ? WHERE id = ?",
		product.Name, product.CategoryId, product.Stock, product.Description, product.UpdatedAt, id,
	)
	if err != nil {
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false
	}

	return rowsAffected > 0
}

// Delete menghapus produk berdasarkan ID
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

// Detail mengambil data produk berdasarkan ID
func Detail(id int) (entities.Product, error) {
	row := config.DB.QueryRow(`
		SELECT 
			products.id,
			products.name,
			products.category_id,
			categories.name AS category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = ?
	`, id)

	var product entities.Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.CategoryId,
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, err
		}
		return product, err
	}

	return product, nil
}
