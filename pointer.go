package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Pass-by-Value
func main() {
	address1 := Address{"Sleman", "Yogyakarta", "Indonesia"}
	address2 := address1

	address2.City = "New-York"
	fmt.Println(address1, address2)

	address3 := Address{"Surabaya", "Jawa-Timur", "Indonesia"}
	address4 := &address3

	//var address3 Address = Address{"Surabaya", "Jawa-Timur", "Indonesia"}
	//var address4 *Address = &address3

	address4.City = "New-York"
	fmt.Println(address3, address4)
}
