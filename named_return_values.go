package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Tom"
	middleName = "Mary" // -> Bisa tidak di declare, dan akan jadi ""
	lastName = "Anderson"
	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
