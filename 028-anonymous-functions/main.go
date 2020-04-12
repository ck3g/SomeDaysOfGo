package main

import "fmt"

func main() {
	// Anonymous function
	func() {
		fmt.Println("Define and run")
	}()

	// func expression
	f := func() {
		fmt.Println("Define, assign to a variable, then run")
	}
	f()
}
