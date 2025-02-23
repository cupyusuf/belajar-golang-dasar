package main

import (
	"Udemy/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Yusuf")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.Version) //  private
	// fmt.Println(helper.sayGoodBye("Yusuf")) // private
}
