package main

import (
	"fmt"
)

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(4,3)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r,im)
}