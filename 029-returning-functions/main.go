package main

import "fmt"

func main() {
	a := bar()

	fmt.Printf("The type of a is %T\n", a)
	fmt.Printf("The value of a() is %v\n", a())
}

func bar() func() int {
	return func() int {
		return 503
	}
}
