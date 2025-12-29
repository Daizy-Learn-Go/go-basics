package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", c.Name)
}

func main() {
	var customer1 Customer
	customer1.Name = "John Doe"
	customer1.Address = "San John Doe"
	customer1.Age = 23
	fmt.Println(customer1)

	customer2 := Customer{
		Name:    "Jane Doe",
		Address: "Yogyakarta",
		Age:     23,
	}
	fmt.Println(customer2)
	customer2.sayHello(customer1.Name)

	customer3 := Customer{"Joko Doe", "Solo", 30}
	fmt.Println(customer3)
}
