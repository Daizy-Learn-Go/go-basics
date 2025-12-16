package main

import "fmt"

func getGoodBye(name string) string {
	return "GoodBye " + name
}

func main() {
	example := getGoodBye
	tatoeba := getGoodBye

	fmt.Println(example("Tod"))
	fmt.Println(tatoeba("Max"))
}
