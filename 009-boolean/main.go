package main

import "fmt"

var x bool

func main() {
	fmt.Println("x is", x)
	x = true
	fmt.Println("now x is", x)

	a := 7
	b := 42
	fmt.Println("a == b is", a == b)
	fmt.Println("a != b is", a != b)
	fmt.Println("a < b is", a < b)
	fmt.Println("a > b is", a > b)
}