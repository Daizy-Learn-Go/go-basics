package main

import "fmt"

func random() any {
	//return "OK"
	return 100
}

func main() {
	var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	switch result.(type) {
	case string:
		fmt.Println("String", result.(string))
	case int:
		fmt.Println("Int", result.(int))
	default:
		fmt.Println("Unknow Type")
	}
}
