package main

import "fmt"

func main() {
	c := make(chan int)

	// Populate the channel
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // Close the channel
	}()

	// Read everything what's in the channel
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
