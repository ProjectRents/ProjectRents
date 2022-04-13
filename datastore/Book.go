package datastore

import (
	"fmt"
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

func (c *BookDB) GetAllDataBook() []entities.Book {
	// Ambil semua data

	var books []entities.Book

	if err := c.DB.Find(&books).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat get data book", err)
	}

	return books
}

func (c BookDB) EditBook(title, isbn, author string, user_id, book_id uint) (string) {
	// Edit data
	
	trx := c.DB.Where("user_id = ? AND id = ?", user_id, book_id).Model(&entities.Book{}).Updates(entities.Book{
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if trx.RowsAffected == 0 {
		return "GAGAL MENGEDIT BUKU"
	}
	
	return "BERHASIL MENGUPDATE BUKU"
}

func (c BookDB) DeleteBook(user_id, book_id uint) (string) {
	// Hapus data
	res := []entities.Book{}

	trx := c.DB.Where("user_id = ? AND id = ?", user_id, book_id).Delete(&res)
	
	if trx.RowsAffected == 0 {
		return "GAGAL MENGHAPUS BUKU"
	}

	return "BERHASIL MENGHAPUS BUKU"
}
