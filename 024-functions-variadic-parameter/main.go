package main

import "fmt"

func main() {
	foo()
	foo(1)
	foo(1, 2, 3)

	a := []int{1, 2, 3, 4}
	foo(a...) // it has to be a final param
}

func foo(x ...int) {
	fmt.Println("what is x?", x)
	fmt.Printf("What type is x? %T\n", x)
	fmt.Println()
}
