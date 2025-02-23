package main

import "fmt"

type Address struct {
	city, province, country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.country = "Indonesia"
}

func main() {
	address := Address{"Cimahi", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)
	fmt.Println(address)
}
