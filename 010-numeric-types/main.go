package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("x = %v of type '%T'\n", x, x)

	y := 42.4242
	fmt.Printf("y = %v of type '%T'\n", y, y)

	var z float64
	z = 10
	fmt.Printf("z = %v of type '%T'\n", z, z)
}
