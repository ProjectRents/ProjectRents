package rent

import (
	"fmt"
)

func InputIDRent(user uint) (user_id, rent_id uint) {
	// Input data
	user_id = user

	fmt.Print("Masukkan ID Rent\t: ")
	fmt.Scanln(&rent_id)

	return 
}