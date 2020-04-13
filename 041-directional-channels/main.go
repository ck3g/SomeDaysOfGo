package main

import (
	"fmt"
)

func main() {
	// Channel
	c := make(chan int, 2)

	c <- 503
	c <- 603

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Println()

	// Send-only channel
	c2 := make(chan<- int, 2)

	c2 <- 503
	c2 <- 603

	// invalid operation: <-c2 (receive from send-only type chan<- int)
	// fmt.Println(<-c2)
	// fmt.Println(<-c2)
	fmt.Printf("%T\n", c2)
	fmt.Println()

	// Receive-only channel
	c3 := make(<-chan int, 2)

	// c3 <- 503 // invalid operation: c3 <- 503 (send to receive-only type <-chan int)
	// c3 <- 603 // invalid operation: c3 <- 603 (send to receive-only type <-chan int)

	// fmt.Println(<-c3) // Nothing to read: all goroutines are asleep - deadlock!
	// fmt.Println(<-c3) // Nothing to read: all goroutines are asleep - deadlock!
	fmt.Printf("%T\n", c3)
	fmt.Println()
}
