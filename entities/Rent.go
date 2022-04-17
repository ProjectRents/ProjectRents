package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserBook struct {
	gorm.Model
	UserID      uint
	BookID      uint `gorm:"primarykey"`
	ReturnDate	time.Time
}
