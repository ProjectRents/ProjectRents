package datastore

import (
	"project_apps/entities"
	"strings"
)

func (connect *DatabaseDB) CreateBook(title string, isbn string, author string, id uint) (string) {
	// Insert data
	
	result := connect.DB.Create(&entities.Book{
		UserID: id,
		Title:  strings.ToUpper(title),
		Isbn:   strings.ToUpper(isbn),
		Author: strings.ToUpper(author),
	})

	if result.Error != nil {
		return "FAILED, PLEASE TRY AGAIN"
	}

	return "SUCCESS, BOOK ADDED"
}

func (connect *DatabaseDB) GetBook(user_id uint) ([]entities.Book, string) {
	// Get data
	search := []entities.Book{}
	connect.DB.Where("user_id = ?", user_id).Find(&search)

	if len(search) == 0 {
		return nil, "FAILED, BOOKS NOT FOUND"
	}

	return search, "SUCCESS, DATA DITAMPILKAN"
}

func (connect DatabaseDB) GetAllDataBook() ([]entities.Book, string) {
	// Ambil semua data

	var books []entities.Book
	connect.DB.Find(&books)

	if len(books) == 0 {
		return nil, "FAILED, BOOKS NOT FOUND"
	}

	return books, "SUCCESS, DATA DITAMPILKAN"
}

func (connect DatabaseDB) EditBook(title, isbn, author string, user_id, book_id uint) (string) {
	// Edit data

	if title != "" || isbn != "" || author != ""  {
		result := connect.DB.Where("user_id = ? AND id = ?", user_id, book_id).Updates(entities.Book{
			Title:  strings.ToUpper(title),
			Isbn:   strings.ToUpper(isbn),
			Author: strings.ToUpper(author),
		})

		if result.RowsAffected == 0 {
			return "FAILED, PLEASE TRY AGAIN"
		}
	} else {
		return "NOTHING TO EDIT"
	}

	return "SUCCESS UPDATE BOOK"
}

func (connect DatabaseDB) DeleteBook(user_id, book_id uint) (string) {
	// Hapus data

	search := entities.Book{}
	result := connect.DB.Where("user_id = ? AND id = ?", user_id, book_id).Delete(&search)
	
	if result.RowsAffected == 0 {
		return "FAILED, DELETE BOOK"
	}

	return "SUCCESS, DELETE BOOK"
}