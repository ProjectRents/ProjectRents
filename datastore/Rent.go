package datastore

import (
	"project_apps/entities"
	"time"
)

func (connect DatabaseDB) CreateRent(user_id, book_id uint, return_date time.Time) string {
	// Insert data

	search := entities.Rent{}
	connect.DB.Where("book_id = ?", book_id).Find(&search)

	if search.BookID == book_id {
		return "FAILED, BOOKS ALREADY RENTED"
	} else {
		result := connect.DB.Create(&entities.Rent{
			UserID:     user_id,
			BookID:     book_id,
			ReturnDate:	return_date,
		})
	
		if result.Error != nil {
			return "FAILED, BOOKS NOT FOUND"
		}
	}

	return "SUCCESS RENT"
}

func (connect DatabaseDB) ReturnRent(user_id, book_id uint) string {
	// Hapus data

	search := entities.Rent{}
	result := connect.DB.Where("user_id = ? AND book_id = ?", user_id, book_id).Delete(&search)

	if result.RowsAffected == 0 {
		return "FAILED, BOOKS NOT FOUND"
	}

	return "SUCCESS RETURN"
}

func (connect DatabaseDB) GetRent(user_id uint) ([]entities.Rent, string) {
	// Get data
	
	search := []entities.Rent{}
	connect.DB.Where("user_id = ?", user_id).Find(&search)

	if len(search) == 0 {
		return nil, "FAILED, RENT NOT FOUND"
	}

	return search, "SUCCESS, DATA DITAMPILKAN"
}
