package entities

import (
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID      uint
	BookID      uint `gorm:"primarykey"`
	Return_date string
}
