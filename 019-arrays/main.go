package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)

	x[3] = 42
	fmt.Println(x)

	// returns the length of an array
	fmt.Println(len(x))
}
