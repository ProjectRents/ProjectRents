package datastore

import (
	"project_apps/entities"
	"strings"
)

func (connect *DatabaseDB) GetAllDataUser(user_id uint) ([]entities.User) {
	// Get data

	users := []entities.User{}
	connect.DB.Where("id = ?", user_id).Find(&users)

	return users
}

func (connect DatabaseDB) CreateUser(name string, email string, password string, alamat string) string {
	// Insert data

	search := entities.User{}
	connect.DB.Where("email = ?", email).Find(&search)
	
	if search.Email == email {
		return "EMAIL ALREADY EXIST, PLEASE CHANGE EMAIL"
	} else {
		result := connect.DB.Create(&entities.User{
			Name:    name,
			Email:   email,
			Password: password,
			Alamat:   alamat,
		})

		if result.RowsAffected == 0 {
			return "FAILED, PLEASE REGISTER AGAIN"
		}
	}

	return "SUCCESS REGISTER, PLEASE LOGIN!"
}

func (connect DatabaseDB) Login(email string, password string) (uint, string) {
	// Login
	
	search := entities.User{}
	connect.DB.Where("email = ? AND password = ?", email, password).Find(&search)

	if search.Email != email || search.Password != password {
		return 0, "FAILED, EMAIL OR PASSWORD IS WRONG"
	}

	name := "HELLO, " + strings.ToUpper(search.Name)
	return search.ID, name
}

func (connect DatabaseDB) EditUser(name, email, password, alamat string, user_id uint) (string) {
	// Edit data

	if name != "" || email != "" || password != "" || alamat != "" {
		search := entities.User{}
		connect.DB.Where("email = ?", email).Find(&search)

		if search.Email != email {
			result := connect.DB.Where("id = ?", user_id).Updates(entities.User{
				Name:    name,
				Email:   email,
				Password: password,
				Alamat:   alamat,
			})
	
			if result.RowsAffected == 0 {
				return "FAILED, PLEASE TRY AGAIN"
			}
		} else {
			return "EMAIL ALREADY EXIST, PLEASE CHANGE EMAIL"
		}
	} else {
		return "NOTHING TO EDIT"
	}

	return "SUCCESS UPDATE USER"
}

func (connect DatabaseDB) DeleteUser(user_id uint) (string) {
	// Hapus data
	search := entities.User{}
	result := connect.DB.Where("id = ?", user_id).Delete(&search)
	
	if result.RowsAffected == 0 {
		return "FAILED, PLEASE TRY AGAIN"
	}

	return "SUCCESS DELETE USER"
}
