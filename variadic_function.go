package main

import "fmt"

func sumAll(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10))

	fmt.Println("\n--== Slice Param to VarArgs ==--")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumAll(numbers...))
}
