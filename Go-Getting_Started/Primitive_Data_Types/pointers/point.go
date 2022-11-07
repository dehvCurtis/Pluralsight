package main

import (
	"fmt"
)

func main() {
	// var firstName *string = new(string)
	// *firstName = "Arthur"
	// fmt.Println(*firstName)

	firstName := "arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Trish"
	fmt.Println(ptr, *ptr)
}