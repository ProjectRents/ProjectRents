package datastore

import "gorm.io/gorm"

type DatabaseDB struct {
	DB *gorm.DB
}