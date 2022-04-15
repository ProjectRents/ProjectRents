package helper

import (
	"fmt"
)


func SetTitleMenu(input string) {
	switch input{
	case "1":
		Separator()
		fmt.Println("ADD NEW ACCOUNT")
		Separator()
	case "2":
		Separator()
		fmt.Println("MENU LOGIN")
		Separator()
	case "3":
		Separator()
		fmt.Println("BOOKS LIST")
	case "99":
		Separator()
		fmt.Println("Quitting...")
		fmt.Println("TERIMA KASIH TELAH MENGGUNAKAN APLIKASI INI")
	}
}


func SetTitleLogin(input string) {
	switch input {
	case "2":
		Separator()
		fmt.Println("MY BOOK LIST")
	case "3":
		Separator()
		fmt.Println("ADD BOOK")
		Separator()
	case "4":
		Separator()
		fmt.Println("EDIT BOOK")
		Separator()
	case "5":
		Separator()
		fmt.Println("DELETE BOOK")
		Separator()
	case "6":
		Separator()
		fmt.Println("RENT BOOK")
		Separator()
	case "7":
		Separator()
		fmt.Println("RENT BOOK LIST")
	case "8":
		Separator()
		fmt.Println("RETURN RENT BOOK")
	}
}

func SetTitleProfile(input string) {
	switch input{
	case "1":
		Separator()
		fmt.Println("MY PROFILE")
	case "2":
		Separator()
		fmt.Println("CHANGE DATA PROFILE")
		Separator()
	case "3":
		Separator()
		fmt.Println("DELETE ACCOUNT")
	}
}