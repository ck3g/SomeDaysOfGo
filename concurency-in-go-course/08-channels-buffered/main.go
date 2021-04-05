package main

import "fmt"

func main() {
	ch := make(chan int, 6) // set the capacity to 6
	// Use unbuffered channel to see the difference
	// ch := make(chan int)

	go func() {
		defer close(ch)

		// TODO: send all interator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
