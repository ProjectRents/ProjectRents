package utilities

import (
	"bufio"
	"fmt"
	"os"
	"project_apps/config"
	"project_apps/datastore"
	"project_apps/entities"
	"project_apps/utilities/book"
	"project_apps/utilities/helper"
	"project_apps/utilities/rent"
	"project_apps/utilities/user"
)

func MenuUtama() {
	// Get database connection
	connection := config.ReadEnv()
	db := config.Database(connection)

	// Get Akses
	Akses := datastore.DatabaseDB{DB: db}
	
	// Migrate tables
	db.AutoMigrate(&entities.User{}, &entities.Book{}, &entities.Rent{})

	if db.Error != nil {
		fmt.Println(db.Error)
		return
	}

	// Menu
	helper.List()
	Menu := bufio.NewScanner(os.Stdin)
	for Menu.Scan() {
		line := Menu.Text()
		if line== "99" {
			helper.SetTitleMenu(line)
			break
		}

		switch line {
		case "1":
			// Add Account
			
			helper.SetTitleMenu(line)
			status := Akses.CreateUser(user.InputUser())
			helper.SetStatus(status)
		case "2":
			// Login
			
			helper.SetTitleMenu(line)
			user_id, name := Akses.Login(user.InputLogin())

			if user_id != 0 {
				helper.SetStatus(name)
			} else {
				helper.SetStatus(name)
				break
			}

			helper.ListLogin()
			Login := bufio.NewScanner(os.Stdin)
			for Login.Scan() {
				line := Login.Text()
				if line == "99" {
					fmt.Println("Logout...")
					break
				}

				switch line {
				case "1":
					// Profile
					
					helper.SetTitleProfile(line)
					helper.ListUser()
					User := bufio.NewScanner(os.Stdin)
					for User.Scan() {
						line := User.Text()
						if line == "99" {
							fmt.Println("BACK...")
							break
						}

						switch line {
						case "1":
							// Get Data User

							helper.SetTitleProfile(line)
							users := Akses.GetAllDataUser(user_id)

							for _, user := range users {
								helper.Separator()
								fmt.Println("Nama\t:", user.Name)
								fmt.Println("Email\t:", user.Email)
								fmt.Println("Alamat\t:", user.Alamat)
							}
						case "2":
							// Edit Profile

							helper.SetTitleProfile(line)
							status := Akses.EditUser(user.InputUbahUser(user_id))
							helper.SetStatus(status)
						case "3":
							// Delete Account

							helper.SetTitleProfile(line)
							status := Akses.DeleteUser(user_id)
							helper.SetStatus(status)
							helper.List()
							Menu.Scan()
						default:
							helper.Separator()
							fmt.Println("MENU NOT FOUND")
						}

						helper.ListUser()
					}
				case "2":
					// Get Book

					helper.SetTitleLogin(line)
					Users, status := Akses.GetBook(user_id)
					helper.SetStatus(status)

					for _, user := range Users {
						helper.Separator()
						fmt.Println("ID\t\t:", user.ID)
						fmt.Println("Judul\t\t:", user.Title)
						fmt.Println("ISBN\t\t:", user.Isbn)
						fmt.Println("Pengarang\t:", user.Author)
					}
				case "3":
					// Add Book

					helper.SetTitleLogin(line)
					status := Akses.CreateBook(book.InputBook(user_id))
					helper.SetStatus(status)
				case "4":
					// Edit Book

					helper.SetTitleLogin(line)
					status := Akses.EditBook(book.InputUbahBook(user_id))
					helper.SetStatus(status)
				case "5":
					// Delete Book

					helper.SetTitleLogin(line)
					status := Akses.DeleteBook(book.InputIDBook(user_id))
					helper.SetStatus(status)
				case "6":
					// Add Book

					helper.SetTitleLogin(line)
					status := Akses.CreateRent(rent.InputRent(user_id))
					helper.SetStatus(status)
				case "7":
					// Get Rent
					
					helper.SetTitleLogin(line)
					Rents, status := Akses.GetRent(user_id)
					helper.SetStatus(status)

					for _, rent := range Rents {
						helper.Separator()
						fmt.Println("ID\t\t  :", rent.ID)
						fmt.Println("Tanggal Kembali\t  :", rent.ReturnDate)
						Books, _ := Akses.GetBook(user_id)
						for _, book := range Books {
							fmt.Println("\tID Buku\t  :", book.ID)
							fmt.Println("\tJudul\t  :", book.Title)
							fmt.Println("\tISBN\t  :", book.Isbn)
							fmt.Println("\tPengarang :", book.Author)
						}
					}

				case "8":
					// Delete Rent

					helper.SetTitleLogin(line)
					status := Akses.ReturnRent(book.InputIDBook(user_id))
					helper.SetStatus(status)
				default:
					helper.Separator()
					fmt.Println("MENU NOT FOUND")
				}

				helper.SetStatus(name)
				helper.ListLogin()	
			}
		case "3":
			// List Books

			books, status := Akses.GetAllDataBook()
			helper.SetTitleMenu(line)
			helper.SetStatus(status)

			for _, book := range books {
				helper.Separator()
				fmt.Println("ID\t:", book.ID)
				fmt.Println("Judul\t:", book.Title)
				fmt.Println("Pengarang\t:", book.Author)
				helper.Separator()
			}
		default:
			helper.Separator()
			fmt.Println("MENU NOT FOUND")
		}

		helper.List()
	}

	if err := Menu.Err(); err != nil {
		fmt.Println("Error encountered:", err)
	}
}
