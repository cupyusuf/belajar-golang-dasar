package main

import "errors"

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagiaj Dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		println("Hasil", hasil)
	} else {
		println("Error", err.Error())
	}
}
