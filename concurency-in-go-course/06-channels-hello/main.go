package main

import "fmt"

func main() {
	// Defines a channel of type int
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c // sends a value to a channel
	}(1, 2)

	// TODO: get the value computed from goroutine
	r := <-ch // reads a value from the channel
	fmt.Printf("computed value %v\n", r)
}
