package datastore

import (
	"fmt"
	"project_apps/entities"

	"gorm.io/gorm"
)

type BookDB struct {
	DB *gorm.DB
}

func (c UserDB) CreateBook(title string, isbn string, author string) {
	// Insert data
	result := c.DB.Create(&entities.Book{
		Title:  title,
		Isbn:   isbn,
		Author: author,
	})

	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
}
