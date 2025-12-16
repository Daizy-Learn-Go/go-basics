package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	//var values = [3]int{1, 2, 3}
	var values = [...]int{
		1, 2, 3, // harus pake koma
	}
	values[0] = 100
	fmt.Println(values)

	fmt.Println(len(values))
}
