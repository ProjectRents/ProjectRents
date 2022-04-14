package datastore

import (
	"fmt"
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

	if result.RowsAffected == 0 {
		return "GAGAL MENAMBAH AKUN"
	}

	return "BERHASIL MENAMBAHKAN AKUN"
}

func (c UserDB) Login(email string, password string) (uint, string) {
	// Login
	results := []entities.User{}
	result := c.DB.Model(&entities.User{}).Where("email = ? AND password = ?", email, password).First(&results)

	if result.Error != nil {
		return 0, "Email atau Password Salah"
	}

	var name string
	var id uint
	for _, user := range results {
		id = user.ID
		name = user.Name
	}

	return id, name
}

func (b *UserDB) GetAllDataUser() ([]entities.User, error) {
	res := []entities.User{}

	if err := b.DB.Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat get data book", err)
		return []entities.User{}, err
	}

	return res, nil

}
