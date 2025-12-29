package main

import "fmt"

type Addross struct {
	City, Province, Country string
}

func ChangeCountryToInedonesia(addross *Addross) {
	addross.Country = "Indonesia"
}

func main() {
	//addross := Addross{}
	//ChangeCountryToInedonesia(&addross)

	addross := new(Addross)
	ChangeCountryToInedonesia(addross)

	fmt.Println(addross)
}
