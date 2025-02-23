package main

import (
	"Udemy/database"
	_ "Udemy/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
