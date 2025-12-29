package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
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
