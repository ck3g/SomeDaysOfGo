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

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()

			v := counter
			runtime.Gosched() // Hey computer, go do something else
			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", counter) // Now it's "100"
}
