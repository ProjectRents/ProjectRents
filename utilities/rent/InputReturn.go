package rent

import (
	"fmt"
)

func InputReturn(user uint) (user_id, book_id uint) {
	// Input data
	user_id = user

	fmt.Print("Masukkan ID Buku\t: ")
	fmt.Scanln(&book_id)

	return
}
