package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			// send interator over a channel
			ch <- i
		}

		close(ch) // closing the channel when we stop sending values
	}()

	// range over channel to receive values
	for v := range ch {
		fmt.Println(v)
	}
}
