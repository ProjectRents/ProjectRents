package utilities

import (
	"bufio"
	"fmt"
	"os"
	"project_apps/config"
	"project_apps/datastore"
	"project_apps/entities"
)

func MenuUtama() {
	// Get database connection
	connection := config.ReadEnv()
	db := config.Database(connection)

	// Migrate tables
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Book{})
	
	

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
				separator()
				fmt.Println("MENAMBAHKAN AKUN BARU")
				separator()
				fmt.Println(Input.CreateUser(InputUser()))
				separator()
			case "2":
				// Login
				Input := datastore.UserDB{DB: db}

				fmt.Println("MENU LOGIN")
				separator()
				id, result := Input.Login(InputLogin())

				if id != 0 {
					fmt.Println("Selamat datang,", result)
				} else {
					fmt.Println(result)
					break
				}
				separator()
				
				listLogin()
				Login := bufio.NewScanner(os.Stdin)
				for Login.Scan(){
					line := Login.Text()
					if line == "99" {
						fmt.Println("Logout...")
						separator()
						break
					}

					switch line {
						case "1":
							fmt.Println("Lihat Profil")
							separator()
						case "2":
							fmt.Println("Daftar Buku Saya")
							separator()
						case "3":
							Input := datastore.BookDB{DB: db}
							separator()
							fmt.Println("MENAMBAHKAN BUKU BARU")
							separator()
							fmt.Println(Input.CreateBook(InputBook(id)))
							separator()
						case "4":
							// Input := datastore.BookDB{DB: db}
							separator()
							fmt.Println("MENGUBAH BUKU")
							separator()
							// fmt.Println(Input.UpdateBook(ChangeBook(id)))
							separator()
						case "5":
							// Input := datastore.BookDB{DB: db}
							separator()
							fmt.Println("MENGHAPUS BUKU")
							separator()
							// fmt.Println(Input.DeleteBook(id))
							separator()
						case "6":
							fmt.Println("Pinjam Buku")
							separator()
						case "7":
							fmt.Println("Kembalikan Buku")
							separator()
						case "99":
							fmt.Println("Logout...")
						default:
							fmt.Println("Menu Tidak Tersedia")
					}

					listLogin()
				}		
		case "3":
			Get := datastore.BookDB{DB: db}
			results := Get.GetAllDataBook()
			for _, result := range results {
				separator()
				fmt.Println("ID :", result.ID)
				fmt.Println("Judul :", result.Title)
				fmt.Println("Pengarang :", result.Author)
				separator()
			}
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
