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

// const (
// 	first = iota + 6
// 	second
// )

const (
	first = iota 
	second
	third
)

const (
	fourth = iota
)

func main() {
	fmt.Println(first, second, third, fourth)
}