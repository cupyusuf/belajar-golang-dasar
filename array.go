package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Yusuf"
	names[1] = "Supriadi"
	names[2] = "Yusuf Supriadi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		85,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
