package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched() // Hey computer, go do something else
			v++
			counter = v

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", counter) // Not "100"
}
