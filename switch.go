package main

import "fmt"

func main() {
	name := "Yusuf"

	switch name {
	case "Yusuf":
		fmt.Println("Hello Yusuf")
	case "Firlana":
		fmt.Println("Hello Firlana")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length == 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
