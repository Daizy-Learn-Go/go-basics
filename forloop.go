package main

import "fmt"

func main() {
	//counter := 1
	//for counter <= 5 {
	//	fmt.Println("Ke-", counter)
	//	counter++
	//}

	for counter := 1; counter <= 3; counter++ {
		fmt.Println("Ke-", counter)
	}
	fmt.Println("Done")
	fmt.Println("--== *** ==--")

	names := []string{"Maulk", "Daffa", "Ardo"}
	//for i := 1; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}
	//for index, name := range names {
	//	fmt.Println(index, name)
	//}
	for _, name := range names {
		fmt.Println(name)
	}
}
