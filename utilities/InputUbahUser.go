package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func InputUbahUser(user uint) (name, email, password, alamat string, user_id uint) {

	user_id = user

	fmt.Print("Masukkan nama: ")
	Name := bufio.NewScanner(os.Stdin)
	Name.Scan()
	name = Name.Text()

	fmt.Print("Masukkan email: ")
	Email := bufio.NewScanner(os.Stdin)
	Email.Scan()
	email = Email.Text()

	fmt.Print("Masukkan password: ")
	Password := bufio.NewScanner(os.Stdin)
	Password.Scan()
	password = Password.Text()

	fmt.Print("Masukkan alamat: ")
	Alamat := bufio.NewScanner(os.Stdin)
	Alamat.Scan()
	alamat = Alamat.Text()
	
	return
}