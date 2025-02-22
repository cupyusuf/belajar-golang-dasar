package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Yusuf", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
