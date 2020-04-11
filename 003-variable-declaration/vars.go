package main

import "fmt"

var x0 = 40

// y0 := 30 // syntax error: non-declaration statement outside function body

var z0 int // Declares a variable of type `int` and set it to Zero value

func main() {
	var x = 42 // variable declaration
	fmt.Println("x =", x)
	x = 99 // reassign variable
	fmt.Println("x =", x)

	y := 43 // short variable declaration
	fmt.Println("y =", y)

	z := "Define " + "expression"
	fmt.Println("z =", z)

	fmt.Println("what is z0?", z0)
}
