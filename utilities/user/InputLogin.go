package utilities_user

import (
	"bufio"
	"fmt"
	"os"
)

func InputLogin() (email, password string) {
	fmt.Print("Masukkan email user: ")
	Email := bufio.NewScanner(os.Stdin)
	Email.Scan()
	email = Email.Text()

	fmt.Print("Masukkan password user: ")
	Password := bufio.NewScanner(os.Stdin)
	Password.Scan()
	password = Password.Text()

	return
}