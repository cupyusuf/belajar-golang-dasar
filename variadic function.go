package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}

	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	number := []int{10, 10, 10, 10, 10}
	total = sumAll(number...)
	fmt.Println(total)
}
