package main

import "fmt"

// These variables have a package scope
var y = 42
var z = "Some string"

func main() {
	fmt.Print("y = ", y)
	fmt.Printf(" of type %T\n", y)

	fmt.Print("z = '", z, "'")
	fmt.Printf(" of type %T\n", z)

	// z = 1 // cannot use 1 (type untyped int) as type string in assignment
}
