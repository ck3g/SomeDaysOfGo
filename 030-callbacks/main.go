package main

import "fmt"

func main() {
	callback := func() {
		fmt.Println("Execute the callback")
	}

	doSomeWork(callback)

	fmt.Println()

	doSomeWork(func() {
		fmt.Println("Execute the inline callback")
	})
}

func doSomeWork(callback func()) {
	fmt.Println("Doing some work")
	callback()
}
