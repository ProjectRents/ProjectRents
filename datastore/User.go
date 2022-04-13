package datastore

import (
	"project_apps/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (c UserDB) CreateUser(name string, email string, password string, alamat string) string {
	// Insert data
	result := c.DB.Create(&entities.User{
		Name:     name,
		Email:    email,
		Password: password,
		Alamat:   alamat,
	})

	if result.Error != nil {
		return "GAGAL MENAMBAH AKUN"
	}

	return "BERHASIL MENAMBAH AKUN"
}