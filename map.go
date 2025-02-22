package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yusuf",
		"address": "Cimahi",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Yusuf Supriadi"
	book["ups"] = "Salah"

	delete(book, "ups")

	fmt.Println(book)
}
