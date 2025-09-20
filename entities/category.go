package entities

import "time"

type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time // Perbaikan: UpdateAt -> UpdatedAt untuk konsistensi
}