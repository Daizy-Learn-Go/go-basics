package main

import "fmt"

func main() {
	type NoKTP string

	var ktpMaul NoKTP = "123123"
	var contoh string = "456456"
	var contohKtp NoKTP = "789789"

	fmt.Println(ktpMaul)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}
