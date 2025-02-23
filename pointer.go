package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Cimahi", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.city = "Bandung"

	// operator & 
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	// operator *
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
