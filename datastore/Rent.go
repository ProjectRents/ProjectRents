package datastore

import (
	"project_apps/entities"

	"gorm.io/gorm"
)

type RentDB struct {
	DB *gorm.DB
}

func (c RentDB) CreateRent(user_id, book_id uint, Return_date string) string {
	// Insert data

	result := c.DB.Create(&entities.Rent{
		UserID:      user_id,
		BookID:      book_id,
		Return_date: Return_date,
	})

	if result.Error != nil {
		return "GAGAL MEMINJAM BUKU"
	}

	return "BERHASIL MEMINJAM BUKU"
}

func (c RentDB) ReturnRent(user_id, book_id uint) string {
	// Hapus data
	res := []entities.Rent{}

	trx := c.DB.Where("user_id = ? AND book_id = ?", user_id, book_id).Delete(&res)

	if trx.RowsAffected == 0 {
		return "GAGAL MENGEMBALIKAN BUKU"
	}

	return "BERHASIL MENGEMBALIKAN BUKU"
}
