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

	// Version 3
	bc := make(chan int, 1) // Buffered channel, allows a 1 value to sit in there

	bc <- 504 // This time it works, because the channel is buffered
	// bc <- 505 // Can't put more values into this channel

	fmt.Println(<-bc)

}
