package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func InputReturn() (UserID, BookID, Return_date string) {
	fmt.Print("Masukkan User: ")
	UserID := bufio.NewScanner(os.Stdin)
	UserID.Scan()
	userID = UserID.Text()

	fmt.Print("Masukkan Buku yang mau dikembalikan: ")
	BookID := bufio.NewScanner(os.Stdin)
	BookID.Scan()
	bookID = BookID.Text()

	fmt.Print("Tanggal masuk: ")
	Return_date := bufio.NewScanner(os.Stdin)
	Return_date.Scan()
	return_date = Return_date.Text()

	return
}
