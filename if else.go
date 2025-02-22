package main

import "fmt"

func main() {
	name := "Yusuf"

	if name == "Yusuf" {
		fmt.Println("Hello, Yusuf")
	} else if name == "Supriadi" {
		fmt.Println("Hello, Supriadi")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
