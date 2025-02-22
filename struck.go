package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var c Customer
	c.Name = "Yusuf"
	c.Address = "Cimahi"
	c.Age = 25

	fmt.Println(c)
	fmt.Println(c.Name)
	fmt.Println(c.Address)
	fmt.Println(c.Age)

	// Struck literal
	c2 := Customer{
		Name:    "Firlana",
		Address: "Jakarta",
		Age:     23,
	}
	fmt.Println(c2)

	c3 := Customer{"Firlana", "Jakarta", 23}
	fmt.Println(c3)
}
