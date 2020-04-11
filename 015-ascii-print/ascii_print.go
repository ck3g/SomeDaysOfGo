package main

import "fmt"

func main() {
	fmt.Println("ASCII printable characters:")

	for i := 32; i <= 126; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}
