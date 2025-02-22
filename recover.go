package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	massage := recover()
	if massage != nil {
		fmt.Println("Terjadi Error", massage)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}

	massage := recover()
	fmt.Println("Terjadi Error", massage)
}

func main() {
	runApp(true)
}
