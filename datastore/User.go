package datastore

import (
	"project_apps/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (c *UserDB) GetAllDataUser(user_id uint) ([]entities.User) {
	// Get data
	res := []entities.User{}

	err := c.DB.Where("id = ?", user_id).Find(&res).Error

	if err != nil {
		return nil
	}

	return res
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

func (c UserDB) EditUser(name, email, password, alamat string, user_id uint) (string) {
	// Edit data
	
	trx := c.DB.Where("id = ?", user_id).Model(&entities.User{}).Updates(entities.User{
		Name:    name,
		Email:   email,
		Password: password,
		Alamat:   alamat,
	})

	if trx.RowsAffected == 0 {
		return "GAGAL MENGEDIT USER"
	}
	
	return "BERHASIL MENGUPDATE USER"
}

func (c UserDB) DeleteUser(user_id uint) (string) {
	// Hapus data
	res := []entities.User{}

	trx := c.DB.Where("id = ?", user_id).Delete(&res)
	
	if trx.RowsAffected == 0 {
		return "GAGAL MENGHAPUS USER"
	}

	return "BERHASIL MENGHAPUS USER"
}
