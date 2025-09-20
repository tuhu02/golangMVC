package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	CategoryId  uint     // Tambahan: ID kategori untuk relasi database
	Category    Category // Struct kategori untuk join query
	Stock       int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}