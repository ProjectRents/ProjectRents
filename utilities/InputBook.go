package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func InputBook() (title, isbn, author string) {
	fmt.Print("Masukkan judul buku: ")
	Title := bufio.NewScanner(os.Stdin)
	Title.Scan()
	title = Title.Text()

	fmt.Print("Masukkan kode ISBN: ")
	Isbn := bufio.NewScanner(os.Stdin)
	Isbn.Scan()
	isbn = Isbn.Text()

	fmt.Print("Masukkan nama pengarang: ")
	Author := bufio.NewScanner(os.Stdin)
	Author.Scan()
	author = Author.Text()
	
	return
}