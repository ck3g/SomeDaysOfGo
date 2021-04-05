package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second) // increase the time to see the timeout message
		ch <- "one"
	}()

	// Imlement timeout for receive on channel ch
	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}
