package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print("Br-", i)
		fmt.Print(", ")
	}
	println()
	for i := 0; i < 10; i++ {
		if i%2 > 0 {
			continue
		}
		fmt.Print("Cont-", i)
		fmt.Print(", ")
	}
}
