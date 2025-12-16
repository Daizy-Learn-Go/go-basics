package main

import "fmt"

func getFullName() (string, string) {
	return "Tom", "Jack"
}

func main() {
	first, second := getFullName()
	fmt.Println(first, second)

	first2, _ := getFullName()
	fmt.Println(first2)
}
