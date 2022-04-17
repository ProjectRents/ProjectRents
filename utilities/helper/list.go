package helper

import "fmt"

func List() {
	Separator()
	fmt.Println("Pilihan Menu :")
	fmt.Println("1\tRegister Account")
	fmt.Println("2\tLogin Account")
	fmt.Println("3\tLihat Daftar Buku")
	fmt.Println("99\tExit")
	fmt.Print("Masukkan 'number' untuk memilih menu : ")
}

func ListLogin() {
	Separator()
	fmt.Println("Pilihan Menu :")
	fmt.Println("1\tProfile")
	fmt.Println("2\tDaftar Buku Saya")
	fmt.Println("3\tTambah Buku")
	fmt.Println("4\tUpdate Buku")
	fmt.Println("5\tHapus Buku")
	fmt.Println("6\tPinjam Buku")
	fmt.Println("7\tDaftar Peminjaman")
	fmt.Println("8\tKembalikan Buku")
	fmt.Println("99\tLogout")
	fmt.Print("Masukkan 'number' untuk memilih menu : ")
}

func ListUser() {
	Separator()
	fmt.Println("Pilihan Menu :")
	fmt.Println("1\tLihat Profile")
	fmt.Println("2\tUpdate Profile")
	fmt.Println("3\tDelete Account")
	fmt.Println("99\tKembali")
	fmt.Print("Masukkan 'number' untuk memilih menu : ")
}

