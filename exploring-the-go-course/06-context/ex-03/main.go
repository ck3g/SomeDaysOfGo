package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	time.Sleep(time.Second) // finished
	// time.Sleep(2 * time.Second) // not finished

	select {
	case <-ctx.Done():
		fmt.Println("not finished")
	default:
		fmt.Println("finished")
	}
}
