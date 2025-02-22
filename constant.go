package main

import "fmt"

func main() {
	const firstName string = "Yusuf"
	const lastName = "Supriadi"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Error
	// firstName = "Tidak bisa diubah"
	// lastName = "Tidak bisa diubah"
}
