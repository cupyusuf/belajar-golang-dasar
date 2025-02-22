package main

import "fmt"

func main() {
	type NoKTP string

	var ktpYusuf = 3277020209990001

	fmt.Println(ktpYusuf)
	fmt.Println(NoKTP("3277020209990002"))
}
