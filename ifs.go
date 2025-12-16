package main

import "fmt"

func main() {
	name := "Awokaowkaow"

	if name == "Maulk" {
		fmt.Println("Hello Maulk")
	} else if name == "Admin" {
		fmt.Println("Hello Admin")
	} else if name == "User" {
		fmt.Println("Hello User")
	} else {
		fmt.Println("Who are you?")
	}

	if len(name) > 5 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("name is good")
	}
}
