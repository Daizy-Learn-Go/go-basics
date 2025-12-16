package main

import "fmt"

func main() {
	//var slice []string
	months := [...]string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October",
		"November", "December",
	}
	slice46 := months[4:6]
	fmt.Print(slice46)
	fmt.Println(" = ", len(slice46), ", ", cap(slice46))

	slice512 := months[5:]
	fmt.Print(slice512)
	fmt.Println(" = ", len(slice512), ", ", cap(slice512))

	slice06 := months[6:]
	fmt.Print(slice06)
	fmt.Println(" = ", len(slice06), ", ", cap(slice06))

	fmt.Println("=============================================")

	days := [...]string{
		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
	}
	fmt.Println(days)
	fmt.Println("\n--== Slice 1 ==--")
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(days)

	fmt.Println("\n--== Append ==--")
	daySlice2 := append(daySlice1, "Holiday")
	daySlice2[0] = "Old Saturday"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("\n--== Make ==--")
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "One"
	newSlice[1] = "Two"
	// newSlice[2] = "Three" -> error, harus pake append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Three")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Awikwok"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("\n--== Copy ==--")
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fmt.Println("\n--== Array vs Slice ==--")
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
