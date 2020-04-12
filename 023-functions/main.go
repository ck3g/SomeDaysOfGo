package main

import "fmt"

func main() {
	foo()
	bar("string")
	returnValue := baz()
	fmt.Println("baz return value is", returnValue)

	status, result := doSomething()
	fmt.Println("doSomething result", status, result)
}

// Syntax
// func (r received) identifier(parameters) (returns(s)) { ... }

func foo() {
	fmt.Println("foo function")
}

func bar(str string) {
	fmt.Println("bar function with argument:", str)
}

func baz() string {
	return "BazString"
}

func doSomething() (string, bool) {
	return "ok", true
}
