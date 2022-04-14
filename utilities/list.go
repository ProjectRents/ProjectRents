package utilities

import "fmt"

func list() {
	fmt.Println("Pilihan Menu :")
	fmt.Println("1. Register Account")
	fmt.Println("2. Login Account")
	fmt.Println("3. Lihat Daftar Buku")
	fmt.Println("99. Exit")
	fmt.Print("Masukkan 'number' untuk memilih menu : ")
}

func listLogin() {
	fmt.Println("Pilihan Menu :")
	fmt.Println("1. Profile")
	fmt.Println("2. Daftar Buku Saya")
	fmt.Println("3. Tambah Buku")
	fmt.Println("4. Update Buku")
	fmt.Println("5. Hapus Buku")
	fmt.Println("6. Pinjam Buku")
	fmt.Println("7. Daftar Peminjaman")
	fmt.Println("8. Kembalikan Buku")
	fmt.Println("99. Logout")
	fmt.Print("Masukkan 'number' untuk memilih menu : ")
}

func listUser() {
	fmt.Println("Pilihan Menu :")
	fmt.Println("1. Lihat Profile")
	fmt.Println("2. Update Profile")
	fmt.Println("3. Delete Account")
	fmt.Println("99. Kembali")
	fmt.Print("Masukkan 'number' untuk memilih menu : ")
}

