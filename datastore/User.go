package datastore

import (
	"fmt"
	"project_apps/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (c UserDB) CreateUser(name string, email string, password string, alamat string) {
	// Insert data
	result := c.DB.Create(&entities.User{
		Name:     name,
		Email:    email,
		Password: password,
		Alamat:   alamat,
	})

	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
}