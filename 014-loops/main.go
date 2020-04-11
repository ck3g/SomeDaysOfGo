package main

import "fmt"

func main() {
	// for initialization; condition; post
	for i := 0; i <= 100; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	for i := 5; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tThe inner loop: %d\n", j)
		}
	}
	fmt.Println()

	a := 1
	for a < 3 {
		a++
		fmt.Printf("a = %v\t", a)
	}
	fmt.Println()

	b := 0
	for {
		b++
		if b > 10 {
			break
		}

		if b%2 != 0 {
			continue
		}

		fmt.Printf("b = %v ", b)
	}
	fmt.Println()
}
