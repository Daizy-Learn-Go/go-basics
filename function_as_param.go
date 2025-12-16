package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println(filter(name))
}

func profanityFilter(name string) string {
	if name == "Asu" {
		return "***"
	} else {
		return name
	}
}

func main() {
	filter := profanityFilter
	//sayHelloWithFilter("Asu", profanityFilter)
	sayHelloWithFilter("Asu", filter)
}
