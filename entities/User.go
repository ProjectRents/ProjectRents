package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(35);not null"`
	Email string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"type:varchar(50);not null"`
	Alamat string `gorm:"type:varchar(35);not null"`
	Books []Book
}