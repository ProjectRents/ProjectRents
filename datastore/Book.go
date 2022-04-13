package datastore

import (
	"project_apps/entities"

	"gorm.io/gorm"
)

type BookDB struct {
	DB *gorm.DB
}

func (c BookDB) CreateBook(title string, isbn string, author string) string {
	// Insert data
	result := c.DB.Create(&entities.Book{
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if result.Error != nil {
		return "GAGAL MENAMBAH BUKU"
	}

	return "BERHASIL MENAMBAH BUKU"
}

func (c BookDB) UpdateBook(title string, isbn string, author string) string {
	// Insert data
	result := c.DB.Create(&entities.Book{
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if result.Error != nil {
		return "GAGAL MENGUBAH BUKU"
	}

	return "BERHASIL MENGUBAH BUKU"
}

func (c BookDB) DeleteBook(title string, isbn string, author string) string {
	// Insert data
	result := c.DB.Create(&entities.Book{
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if result.Error != nil {
		return "GAGAL MENGUBAH BUKU"
	}

	return "BERHASIL MENGUBAH BUKU"
}
