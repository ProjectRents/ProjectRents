package datastore

import (
	"fmt"
	"project_apps/entities"

	"gorm.io/gorm"
)

type BookDB struct {
	DB *gorm.DB
}

func (b *BookDB) GetAllDataBook() []entities.Book {
	results := []entities.Book{}
	if err := b.DB.Find(&results).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat get data book", err)
	}

	return results
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

func (c BookDB) UpdateBook(title string, isbn string, author string, id uint) string {
	// Update data
	result := c.DB.Save(&entities.Book{
		UserID: id,
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if result.Error != nil {
		return "GAGAL MENGUBAH BUKU"
	}

	return "BERHASIL MENGUBAH BUKU"
}

func (c BookDB) DeleteBook(id uint) (string) {
	// Hapus data
	res := []entities.Book{}

	tx := c.DB.Where("id = ?", id).Delete(&res)
	if tx.Error != nil {
		return "GAGAL MENGHAPUS BUKU"
	}
	
	return "BERHASIL MENGHAPUS BUKU"
}
