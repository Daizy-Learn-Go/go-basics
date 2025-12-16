package main

import "fmt"

func main() {
	name := "Maulk"

	switch name {
	case "Maulk":
		fmt.Println("Hello Maulk")
	case "Admin":
		fmt.Println("Hello Admin")
	default:
		fmt.Println("Hello World")
	}

	switch len(name) > 5 {
	case true:
		fmt.Println("Name too long")
	case false:
		fmt.Println("Name is good")
	}

	switch {
	case len(name) > 5:
		fmt.Println("Name too long")
	case len(name) == 0:
		fmt.Println("Name is empty")
	case len(name) < 6:
		fmt.Println("Name is good")
	}
}
