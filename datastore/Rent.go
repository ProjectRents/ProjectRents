package datastore

import (
	"project_apps/entities"
	"time"
)

func (connect DatabaseDB) CreateRent(user_id, book_id uint, return_date time.Time) string {
	// Insert data

	search := entities.UserBook{}
	connect.DB.Where("book_id = ?", book_id).Find(&search)

	if search.BookID == book_id {
		return "FAILED, BOOKS ALREADY RENTED"
	} else {
		result := connect.DB.Create(&entities.UserBook{
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

	search := entities.UserBook{}
	result := connect.DB.Where("user_id = ? AND book_id = ?", user_id, book_id).Delete(&search)

	if result.RowsAffected == 0 {
		return "FAILED, BOOKS NOT FOUND"
	}

	return "SUCCESS RETURN"
}

func (connect DatabaseDB) GetRent(user_id uint) ([]entities.User, []entities.UserBook, string) {
	// Get data
	
	search := []entities.User{}
	connect.DB.Where("id = ?", user_id).Preload("Books").Find(&search)

	return_date := []entities.UserBook{}
	connect.DB.Where("user_id = ?", user_id).Find(&return_date)

	if (len(search) == 0) || (len(return_date) == 0) {
		return nil, nil, "FAILED, RENT NOT FOUND"
	}

	return search, return_date, "SUCCESS, DATA DITAMPILKAN"
}
