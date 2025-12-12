package main

import "fmt"

func main() {
	// Const -> tidak dapat diubah, tetapi tidak masalah jika tidak digunakan

	//const firstName = "Maulana"
	//const lastName = "Ardiansyah"

	const (
		firstName = "Maulana"
		lastName  = "Ardiansyah"
	)

	// error
	// firstName = "Kocak"
	// lastName = "Ganti Nama"

	fmt.Println(firstName, lastName)
}
