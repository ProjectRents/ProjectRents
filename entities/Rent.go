package entities

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID      uint
	BookID      uint `gorm:"primarykey"`
	ReturnDate	time.Time
}
