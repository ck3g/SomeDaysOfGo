package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is true, so that line is printed")
	}

	if false {
		fmt.Println("This line will never be printed")
	}

	fmt.Println()

	if 2 > 2 {
		fmt.Println("2 is greater than 2. Really?")
	} else {
		fmt.Println("2 is not greater than 2")
	}

	if x := 1; x == 2 {
		fmt.Println("x is equal to two")
	} else if x == 3 {
		fmt.Println("x is equal to three")
	} else {
		fmt.Println("x is visible here as well. It equals to", x)
	}
}
