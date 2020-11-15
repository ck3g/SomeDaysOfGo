package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("GOROUTINES RUNNING %d\n", runtime.NumGoroutine())

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println("Launching goroutine", n)
			for {

				select {
				case <-ctx.Done():
					runtime.Goexit()
				default:
					fmt.Printf("Goroutine %d is doing work\n", n)
					time.Sleep(50 * time.Millisecond)
				}
			}

		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING %d\n", runtime.NumGoroutine())
	cancel()

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING %d\n", runtime.NumGoroutine())
}
