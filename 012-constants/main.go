package main

import "fmt"

// Untyped constants
const a = 1
const b = 1.1
const c = "A"

const (
	a2 = 2
	b2 = 2.2
	c2 = "B"
)

// Typed constants
const (
	a3 int = 3
	b3 float64 = 3.3
	c3 string = "C"
)

func main() {
	fmt.Printf("a = %v of type %T\n", a, a)
	fmt.Printf("b = %v of type %T\n", b, b)
	fmt.Printf("c = %v of type %T\n", c, c)
	fmt.Println()

	fmt.Printf("a2 = %v of type %T\n", a2, a2)
	fmt.Printf("b2 = %v of type %T\n", b2, b2)
	fmt.Printf("c2 = %v of type %T\n", c2, c2)
	fmt.Println()

	fmt.Printf("a3 = %v of type %T\n", a3, a3)
	fmt.Printf("b3 = %v of type %T\n", b3, b3)
	fmt.Printf("c3 = %v of type %T\n", c3, c3)
}
