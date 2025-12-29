package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "id is empty"}
	}

	if id != "Eko" {
		return &notFoundError{Message: "id not found"}
	}

	return nil
}

func main() {
	err := SaveData("Eko", nil)
	if err != nil {
		//if validationError, ok := err.(*validationError); ok {
		//	fmt.Println("Validation error:", validationError.Message)
		//} else if notFoundError, ok := err.(*notFoundError); ok {
		//	fmt.Println("Not found error:", notFoundError.Message)
		//} else {
		//	fmt.Println("Unknown Error:", err)
		//}
		switch errType := err.(type) {
		case *validationError:
			fmt.Println("Validation error:", errType.Message)
		case *notFoundError:
			fmt.Println("Not found error:", errType.Message)
		default:
			fmt.Println("Unknown error:", err)
		}
	} else {
		fmt.Println("Berhasil")
	}
}
