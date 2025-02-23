package main

import "fmt"

type validationError struct {
	Massage string
}

func (v *validationError) Error() string {
	return v.Massage
}

type notFoundError struct {
	Massage string
}

func (n *notFoundError) Error() string {
	return n.Massage
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}
	if id != "yusuf" {
		return &notFoundError{"data not found"}
	}

	// ok
	return nil
}

func main() {
	err := SaveData("yusuf", nil)
	if err != nil {
		// terjadi error

		// if else
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error", err.Error())
		// }

		// switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error", finalError.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
