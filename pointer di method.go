package main

import "fmt"

type Man struct {
	name string
}

func (m *Man) Married() {
	m.name = "Mr. " + m.name
}

func main() {
	yusuf := Man{"Yusuf"}
	yusuf.Married()

	fmt.Println(yusuf.name)
}
