package utilities

import (
	"bufio"
	"fmt"
	"os"
	"project_apps/config"
	"project_apps/datastore"
	"project_apps/entities"
	utilities_book "project_apps/utilities/book"
	utilities_rent "project_apps/utilities/rent"
	utilities_user "project_apps/utilities/user"
)

func MenuUtama() {
	// Get database connection
	connection := config.ReadEnv()
	db := config.Database(connection)

	// Migrate tables
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Book{})
	db.AutoMigrate(&entities.Rent{})

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
			fmt.Println(Input.CreateUser(utilities_user.InputUser()))
			separator()
		case "2":
			// Login
			Input := datastore.UserDB{DB: db}
			separator()
			fmt.Println("MENU LOGIN")
			separator()
			user_id, result := Input.Login(utilities_user.InputLogin())

			if user_id != 0 {
				fmt.Println("Selamat datang,", result)
			} else {
				fmt.Println(result)
				break
			}
			separator()

			listLogin()
			Login := bufio.NewScanner(os.Stdin)
			for Login.Scan() {
				line := Login.Text()
				if line == "99" {
					fmt.Println("Logout...")
					separator()
					break
				}

				switch line {
				case "1":
					// Profile
					separator()
					fmt.Println("PROFIL SAYA")
					separator()

					listUser()
					User := bufio.NewScanner(os.Stdin)
					for User.Scan() {
						line := User.Text()
						if line == "99" {
							fmt.Println("Kembali...")
							separator()
							break
						}

						switch line {
						case "1":
							// Get data user

							Get := datastore.UserDB{DB: db}
							getUser := Get.GetAllDataUser(user_id)
							separator()
							fmt.Println("DETAIL PROFIL")

							for _, result := range getUser {
								separator()
								fmt.Println("Nama:", result.Name)
								fmt.Println("Email:", result.Email)
								fmt.Println("Password:", result.Password)
								fmt.Println("Alamat:", result.Alamat)
								separator()
							}

						case "2":
							// Edit Profile

							Input := datastore.UserDB{DB: db}
							separator()
							fmt.Println("UBAH PROFIL SAYA")

							fmt.Println(Input.EditUser(utilities_user.InputUbahUser(user_id)))
							separator()

						case "3":
							// Delete Profile

							Input := datastore.UserDB{DB: db}
							separator()
							fmt.Println("HAPUS AKUN SAYA")

							fmt.Println(Input.DeleteUser(user_id))
							separator()
							list()
							Menu.Scan()

						default:
							fmt.Println("Menu Tidak Tersedia")
						}

						listUser()
					}
				case "2":
					// Get Book

					Input := datastore.BookDB{DB: db}
					separator()
					fmt.Println("MENAMPILKAN SEMUA BUKU")

					DatabyUser := (Input.GetBook(user_id))
					for _, result := range DatabyUser {
						separator()
						fmt.Println("ID :", result.ID)
						fmt.Println("Judul :", result.Title)
						fmt.Println("Pengarang :", result.Author)
						separator()
					}
				case "3":
					// Add Book

					Input := datastore.BookDB{DB: db}
					separator()
					fmt.Println("MENAMBAHKAN BUKU BARU")
					separator()
					fmt.Println(Input.CreateBook(utilities_book.InputBook(user_id)))
					separator()
				case "4":
					// Edit Book

					Input := datastore.BookDB{DB: db}
					separator()
					fmt.Println("MENGUBAH BUKU")
					separator()
					fmt.Println(Input.EditBook(utilities_book.InputUbahBook(user_id)))
					separator()
				case "5":
					// Delete Book

					Input := datastore.BookDB{DB: db}
					separator()
					fmt.Println("MENGHAPUS BUKU")
					separator()
					fmt.Println(Input.DeleteBook(utilities_book.InputIDBook(user_id)))
					separator()
				case "6":
					// Add Book

					Input := datastore.RentDB{DB: db}
					separator()
					fmt.Println("MEMINJAM BUKU")
					separator()
					fmt.Println(Input.CreateRent(utilities_rent.InputRent(user_id)))
					separator()
				case "7":
					Input := datastore.RentDB{DB: db}
					separator()
					fmt.Println("MEMULANGKAN BUKU")
					separator()
					fmt.Println(Input.ReturnRent(utilities_book.InputIDBook(user_id)))
					separator()
				case "99":
					fmt.Println("Logout...")
				default:
					fmt.Println("Menu Tidak Tersedia")
				}

				listLogin()
			}
		case "3":
			// List Books

			Get := datastore.BookDB{DB: db}
			getAll := Get.GetAllDataBook()
			separator()
			fmt.Println("DAFTAR BUKU")

			for _, result := range getAll {
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
