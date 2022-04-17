package book

import (
	"bufio"
	"fmt"
	"os"
)

func InputBook(u uint) (title, isbn, author string, id uint) {
	// Input data

	id = u
	
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