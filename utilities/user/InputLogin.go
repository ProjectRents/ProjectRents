package user

import (
	"bufio"
	"fmt"
	"os"
)

func InputLogin() (email, password string) {
	fmt.Print("Masukkan email user\t: ")
	Email := bufio.NewScanner(os.Stdin)
	Email.Scan()
	email = Email.Text()

	fmt.Print("Masukkan password user\t: ")
	Password := bufio.NewScanner(os.Stdin)
	Password.Scan()
	password = Password.Text()

	return
}