package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Maulk"
	//person["age"] = "23"

	person := map[string]string{
		"name": "John",
		"age":  "25",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)
	//fmt.Println(person["address"])

	fmt.Println("\n--== Map Function ==--")
	book := make(map[string]string)
	book["title"] = "Go Programming Language"
	book["author"] = "Andrew Johnson"
	book["year"] = "2022"
	fmt.Println(book)
	delete(book, "year")
	fmt.Println(book)
}
