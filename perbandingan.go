package main

import "fmt"

func main() {
	var name1 = "Maul"
	var name2 = "Maul"

	var result bool = name1 == name2
	fmt.Println(result)

	var result2 bool = name1 != name2
	fmt.Println(result2)

	var char1 = "b"
	var char2 = "a"

	var result3 bool = char1 > char2
	fmt.Println(result3)
}
