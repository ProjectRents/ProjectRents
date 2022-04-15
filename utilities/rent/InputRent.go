package rent

import (
	"fmt"
	"time"
)

func InputRent(u uint) (user_id, book_id uint, tanggal_kembali time.Time) {
	// Input data
	
	user_id = u
	fmt.Print("Masukkan ID buku yang dipinjam: ")
	fmt.Scanln(&book_id)
	
	var day int
	fmt.Print("Masukkan Jumlah Hari: ")
	fmt.Scanln(&day)
	tanggal_kembali = time.Now().AddDate(0, 0, day)

	return
}
