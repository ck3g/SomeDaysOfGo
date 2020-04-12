package main

import "fmt"

func main() {
	defer foo() // defer runs then then function scope is about to end
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
