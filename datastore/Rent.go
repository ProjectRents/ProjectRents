package datastore

import (
	"project_apps/entities"

	"gorm.io/gorm"
)

type RentDB struct {
	DB *gorm.DB
}

func (c RentDB) CreateBook(Return_date string) string {
	// Insert data

	result := c.DB.Create(&entities.Rent{

		Return_date: Return_date,
	})

	if result.Error != nil {
		return "GAGAL MEMINJAM BUKU"
	}

	return "BERHASIL MEMINJAM BUKU"
}

func (c RentDB) ReturnBook(Return_date string) string {
	// Insert data

	result := c.DB.Create(&entities.Rent{

		Return_date: Return_date,
	})

	if result.Error != nil {
		return "GAGAL MENGEMBALIKAN BUKU"
	}

	return "BERHASIL MENGEMBALIKAN BUKU"
}
