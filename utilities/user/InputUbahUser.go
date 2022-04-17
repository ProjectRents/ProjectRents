package user

import (
	"bufio"
	"fmt"
	"os"
)

func InputUbahUser(user uint) (name, email, password, alamat string, user_id uint) {

	user_id = user

	fmt.Print("Masukkan nama\t\t: ")
	Name := bufio.NewScanner(os.Stdin)
	Name.Scan()
	name = Name.Text()

	fmt.Print("Masukkan email\t\t: ")
	Email := bufio.NewScanner(os.Stdin)
	Email.Scan()
	email = Email.Text()

	fmt.Print("Masukkan password\t: ")
	Password := bufio.NewScanner(os.Stdin)
	Password.Scan()
	password = Password.Text()

	fmt.Print("Masukkan alamat\t\t: ")
	Alamat := bufio.NewScanner(os.Stdin)
	Alamat.Scan()
	alamat = Alamat.Text()
	
	return
}