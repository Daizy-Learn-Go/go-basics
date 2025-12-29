package main

import "fmt"

type Addressss struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Addressss = new(Addressss)
	var alamat2 *Addressss = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1, alamat2)
}
