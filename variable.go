package main

import "fmt"

func main() {
	// Jika tidak digunakan, tidak dapat di run

	//var name string; name = "Maulana Daffa"
	//var name = "Maulana Daffa"
	name := "Maulana Daffa"
	fmt.Println(name)

	name = "Post Maulone"
	fmt.Println(name)

	var (
		firstName  = "Maulana"
		middleName = "Daffa"
		lastName   = "Ardiansyah"
	)
	fmt.Println(firstName, " ", middleName, " ", lastName)
	fmt.Println(firstName, middleName, lastName)
}
