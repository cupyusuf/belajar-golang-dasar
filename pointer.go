package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Cimahi", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.city = "Bandung"

	fmt.Println(address1) // {Cimahi Jawa Barat Indonesia}
	fmt.Println(address2)
}
