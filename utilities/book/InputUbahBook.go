package book

import (
	"bufio"
	"fmt"
	"os"
)

func InputUbahBook(user uint) (title, isbn, author string, user_id, book_id uint) {

	user_id = user

	fmt.Print("Masukkan ID Buku\t: ")
	fmt.Scanln(&book_id)
	
	fmt.Print("Masukkan judul buku\t: ")
	Title := bufio.NewScanner(os.Stdin)
	Title.Scan()
	title = Title.Text()

	fmt.Print("Masukkan kode ISBN\t: ")
	Isbn := bufio.NewScanner(os.Stdin)
	Isbn.Scan()
	isbn = Isbn.Text()

	fmt.Print("Masukkan nama pengarang\t: ")
	Author := bufio.NewScanner(os.Stdin)
	Author.Scan()
	author = Author.Text()

	return
}