package entities

import "gorm.io/gorm"

type DatabaseDB struct {
	DB *gorm.DB
}