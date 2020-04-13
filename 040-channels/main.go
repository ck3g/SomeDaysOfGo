package main

import "fmt"

func main() {
	c := make(chan int)

	// Version 1
	// c <- 503 // fatal error: all goroutines are asleep - deadlock!

	// Version 2
	go func() {
		c <- 503
	}()

	fmt.Println(<-c)
}
