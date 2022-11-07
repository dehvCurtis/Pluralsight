package main

import (
	"fmt"
)

// const (
// 	first = 1
// 	second = "second"
// )

const (
	first = iota
	second = iota
)

func main() {
	fmt.Println(first, second)
}