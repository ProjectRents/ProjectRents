package datastore

import (
	"project_apps/entities"

	"gorm.io/gorm"
)

type BookDB struct {
	DB *gorm.DB
}

func (c BookDB) CreateBook(title string, isbn string, author string, id uint) (string) {
	// Insert data
	
	result := c.DB.Create(&entities.Book{
		UserID: id,
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if result.Error != nil {
		return "GAGAL MENAMBAH BUKU"
	}

	return "BERHASIL MENAMBAH BUKU"
}
