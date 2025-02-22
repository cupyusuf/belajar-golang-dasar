package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Yusuf"
	lastName = "Supriadi"

	return firstName, lastName
}

func main() {
	firstName, lastName := getCompleteName()
	fmt.Println(firstName, lastName)
}
