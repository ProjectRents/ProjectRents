package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func InputUser() (name, email, password, alamat string) {
	fmt.Print("Masukkan nama user: ")
	Name := bufio.NewScanner(os.Stdin)
	Name.Scan()
	name = Name.Text()

	fmt.Print("Masukkan email user: ")
	Email := bufio.NewScanner(os.Stdin)
	Email.Scan()
	email = Email.Text()

	fmt.Print("Masukkan password user: ")
	Password := bufio.NewScanner(os.Stdin)
	Password.Scan()
	password = Password.Text()

	fmt.Print("Masukkan alamat: ")
	Alamat := bufio.NewScanner(os.Stdin)
	Alamat.Scan()
	alamat = Alamat.Text()

	return
}