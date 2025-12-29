package main

import "fmt"

func logging() {
	fmt.Println("Function Done")
}

func runFunc() {
	defer logging()
	fmt.Println("Function Running")
	// jika error, kode dibawah tidak jalan
	// logging()
}

func main() {

}
