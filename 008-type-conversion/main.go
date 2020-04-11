package main

import "fmt"

var a int

type intClone int

var b intClone

func main() {
	a = 1
	fmt.Printf("a = %v of type %T\n", a, a)

	b = 2
	fmt.Printf("b = %v of type %T\n", b, b)

	// a = b // cannot use b (type intClone) as type int in assignment

	a = int(b) // Type conversion
	fmt.Printf("a = %v of type %T\n", a, a)
}
