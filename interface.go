package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println(value.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

func main() {
	person1 := Person{Name: "John"}
	SayHello(person1)

	animal1 := Animal{Name: "Animal"}
	SayHello(animal1)
}
