package main

import "fmt"

func logging() {
	fmt.Println("Selesai memaggil logging")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
