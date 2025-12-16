package main

import "fmt"

func main() {
	var (
		a = 10
		b = 15
		c = a + b
		d = b * a
		e = b / 5
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("--------- += ---------")

	var i = 10
	fmt.Println(i)

	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	fmt.Println("--------- ++ ---------")

	var j = 10
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}
