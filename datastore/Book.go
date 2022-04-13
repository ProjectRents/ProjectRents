package datastore

import (
	"fmt"
	"project_apps/entities"

	"gorm.io/gorm"
)

type BookDB struct {
	DB *gorm.DB
}

func (b *BookDB) GetAllDataBook() ([]entities.Book, error) {
	res := []entities.Book{}

	if err := b.DB.Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat get data book", err)
		return []entities.Book{}, err
	}

	return res, nil
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
	result := c.DB.Save(&entities.Book{
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if result.Error != nil {
		return "GAGAL MENGUBAH BUKU"
	}

	return "BERHASIL MENGUBAH BUKU"
}

func (c BookDB) DeleteBook() ([]entities.Book, error) {
	res := []entities.Book{}

	tx := c.DB.Where("name=?", 3).Delete(BookDB)
	if tx.Error != nil {
		fmt.Println("Delete gagal", tx.Error)
	}
	return res, nil
}
