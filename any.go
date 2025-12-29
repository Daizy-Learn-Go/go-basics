package main

import "fmt"

func Ups() interface{} {
	//return 1
	//return true
	return "Any"
}

func main() {
	var kosong = Ups()
	fmt.Println(kosong)
}
