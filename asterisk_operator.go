package main

import "fmt"

type Addresss struct {
	City, Province, Country string
}

func main() {
	address1 := Addresss{"Surabaya", "Jawa-Timur", "Indonesia"}
	address2 := &address1
	address2.City = "New-York"
	fmt.Println(address1, address2)

	//address2 = &Addresss{"Sleman", "Yogyakarta", "Indonesia"}
	//fmt.Println(address1, address2)

	*address2 = Addresss{"Jakarta", "DKI-Jakarta", "Indonesia"}
	fmt.Println(address1, address2)
}
