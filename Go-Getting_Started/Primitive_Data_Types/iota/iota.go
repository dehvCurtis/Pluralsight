package main

import (
	"fmt"
)

// const (
// 	first = 1
// 	second = "second"
// )

// const (
// 	first = iota
// 	second = iota
// )

const (
	first = iota + 6
	second = 2 << iota
)

func main() {
	fmt.Println(first, second)
}