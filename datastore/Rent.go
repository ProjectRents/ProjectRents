package datastore

import (
	"project_apps/entities"

	"gorm.io/gorm"
)

type RentDB struct {
	DB *gorm.DB
}

func (c RentDB) CreateRent() string {
	// Insert data

	result := c.DB.Save(&entities.Rent{
		UserID:      1,
		BookID:      1,
		Return_date: "2022-10-10",
	})

	if result.Error != nil {
		return "GAGAL MEMINJAM BUKU"
	}

	return "BERHASIL MEMINJAM BUKU"
}

func (c RentDB) ReturnRent() string {
	// Insert data

	result := c.DB.Create(&entities.Rent{
		UserID:      1,
		BookID:      2,
		Return_date: "2022-11-11",
	})

	if result.Error != nil {
		return "GAGAL MENGEMBALIKAN BUKU"
	}

	return "BERHASIL MENGEMBALIKAN BUKU"
}
