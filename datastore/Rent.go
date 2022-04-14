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

	rents := []entities.Rent{}
	c.DB.Where("book_id = ?", book_id).Find(&rents)

	for _, rent := range rents {
		if rent.BookID != book_id {
			result := c.DB.Create(&entities.Rent{
				UserID:      user_id,
				BookID:      book_id,
				Return_date: Return_date,
			})
	
			if result.Error != nil {
				return "GAGAL MEMINJAM BUKU"
			}
		} else {
			return "BUKU SUDAH DI PINJAM"
		}
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

func (c RentDB) GetRent(user_id uint) []entities.Rent {
	// Get data
	rents := []entities.Rent{}

	c.DB.Where("user_id = ?", user_id).Find(&rents)

	return rents
}
