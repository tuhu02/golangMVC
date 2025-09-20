package categorymodel

import (
	"belajar/golangMVC/entities"
	"belajar/golangMVC/config"
	"database/sql"
)

// GetAll mengambil semua data kategori dari database
func GetAll() []entities.Category {
	rows, err := config.DB.Query("SELECT id, name, created_at, updated_at FROM categories")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next(){
		var category entities.Category
		// Perbaikan: UpdateAt -> UpdatedAt
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}
	
	return categories
}

// Create menambahkan kategori baru ke database
func Create(category entities.Category) bool {
	result, err := config.DB.Exec("INSERT INTO categories(name,created_at,updated_at) VALUES (?,?,?)", category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return false
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return false
	}

	return lastInsertId > 0
}

// Delete menghapus kategori berdasarkan ID
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}

// Update mengubah data kategori berdasarkan ID
func Update(id int, category entities.Category) bool {
	result, err := config.DB.Exec(
		"UPDATE categories SET name = ?, updated_at = ? WHERE id = ?", 
		category.Name, category.UpdatedAt, id,
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


// Detail mengambil data kategori berdasarkan ID
func Detail(id int) (entities.Category, error) {
	row := config.DB.QueryRow("SELECT id, name, created_at, updated_at FROM categories WHERE id = ?", id)

	var category entities.Category
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return category, err
		}
		return category, err 
	}

	return category, nil 
}
