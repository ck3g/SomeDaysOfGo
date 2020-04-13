package main

import "fmt"

func main() {
	a := 503
	fmt.Printf("The value of 'a' is %v\n", a)
	fmt.Printf("The address of 'a' is %v\n", &a)
	fmt.Printf("The type of 'a' is %T\n", a)
	fmt.Printf("The type of '&a' is %T\n", &a)
	fmt.Println()

	b := &a // Now b has the same address as a
	fmt.Printf("The address that 'b' points at is %v\n", b)
	fmt.Printf("The value in the address of '*b' is %v\n", *b)
	fmt.Println()

	a = 505 // a and *b share the same address
	fmt.Printf("The new value of 'a' is %v\n", a)
	fmt.Printf("The new value of '*b' is %v\n", *b)
	fmt.Println()

	*b = 404 // a and *b share the same address
	fmt.Printf("The new value of 'a' is %v\n", a)
	fmt.Printf("The new value of '*b' is %v\n", *b)
	fmt.Println()
}
