package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)

			runtime.Gosched() // Hey computer, go do something else
			fmt.Println("atomic counter\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", counter) // Not "100"
}
