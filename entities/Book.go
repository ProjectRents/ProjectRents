package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	UserID 	uint
	Title  	string `gorm:"type:varchar(50);not null"`
	Isbn   	string `gorm:"type:varchar(50);not null"`
	Author 	string `gorm:"type:varchar(50);not null"`
	Users	[]User `gorm:"many2many:user_books"`
}
