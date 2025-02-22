package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

// Contoh Salah
// func runApp(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("APLIKASI ERROR")
// 	}
// 	massage := recover()
// 	fmt.Println("Terjadi error", massage)
// }

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

}

func main() {
	runApp(false)
}
