package main

import "fmt"

var a int
var b string
var c bool

func main() {
	fmt.Println("Unassigned varialbes has a Zero value:")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b, " <- (empty string)")
	fmt.Println("c = ", c)
}
