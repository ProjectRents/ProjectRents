package utilities_rent

import (
	"bufio"
	"fmt"
	"os"
)

func InputRent(u uint) (user_id, book_id uint, tanggal_kembali string) {
	// Input data
	user_id = u
	fmt.Print("Masukkan ID buku yang mau pinjam: ")
	fmt.Scanln(&book_id)

	fmt.Print("Tanggal Kembali: ")
	Kalender := bufio.NewScanner(os.Stdin)
	Kalender.Scan()
	tanggal_kembali = Kalender.Text()

	return
}
