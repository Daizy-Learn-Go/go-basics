package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Error:", message)

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Oops, Error")
	}
	// endApp() dibawah tidak akan dipanggil jika error = true
	// endApp()
}

func main() {
	runApp(false)
	runApp(true)
	fmt.Println("Main Done")
}
