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

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat1.country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
