package main

import "fmt"

func getFullName() (string, string) {
	return "Yusuf", "Supriadi"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	// Menghiraukan Return Value
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
