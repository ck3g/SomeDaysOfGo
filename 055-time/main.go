package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("About to tick. Current time is %v\n", time.Now())
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	go func() {
		fmt.Printf("Before a tick %v\n", time.Now())

		for range ticker.C {
			fmt.Printf("Ticking %v\n", time.Now())
		}
	}()

	fmt.Println("Ending programm in 10 seconds...")
	time.Sleep(10 * time.Second)
}
