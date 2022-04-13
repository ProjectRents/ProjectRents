package utilities

import (
	"bufio"
	"fmt"
	"os"
	"project_apps/config"
	"project_apps/datastore"
)

func MenuUtama() {
	// Get database connection
	connection := config.ReadEnv()
	db := config.Database(connection)
	
	// Migrate tables
	// db.AutoMigrate(&entities.User{})

	if db.Error != nil {
		fmt.Println(db.Error)
		return
	}

	// Menu
	list()
	Menu := bufio.NewScanner(os.Stdin)
	for Menu.Scan() {
		line := Menu.Text()
		if line == "99" {
			fmt.Println("Quitting...")
			break
		}

		switch line {
			case "1":
				// Add Account
				Input := datastore.UserDB{DB: db}

				fmt.Println("MENAMBAHKAN AKUN BARU")
				separator()
				Input.CreateUser(InputUser())
				separator()
				fmt.Println("Selamat anda telah terdaftar")
				separator()
			case "2":
				fmt.Println("Login Account")
			case "3":
				fmt.Println("Lihat Daftar Buku")
			case "99":
				fmt.Println("Exit")
			default:
				fmt.Println("Menu Tidak Tersedia")
				separator()
		}

		list()
	}
		
	if err := Menu.Err(); err != nil {
		fmt.Println("Error encountered:", err)
	}	
}