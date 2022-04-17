package user

import (
	"bufio"
	"fmt"
	"os"
)

func InputUser() (name, email, password, alamat string) {
	fmt.Print("Masukkan nama user\t: ")
	Name := bufio.NewScanner(os.Stdin)
	Name.Scan()
	name = Name.Text()

	fmt.Print("Masukkan email user\t: ")
	Email := bufio.NewScanner(os.Stdin)
	Email.Scan()
	email = Email.Text()

	fmt.Print("Masukkan password user\t: ")
	Password := bufio.NewScanner(os.Stdin)
	Password.Scan()
	password = Password.Text()

	fmt.Print("Masukkan alamat\t\t: ")
	Alamat := bufio.NewScanner(os.Stdin)
	Alamat.Scan()
	alamat = Alamat.Text()

	return
}