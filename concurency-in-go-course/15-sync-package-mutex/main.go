package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// without Lock/Unlock the blance amount is going to be unpredictable

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	withdrawal := func(amount int) {
		mu.Lock()
		balance -= amount
		mu.Unlock()
	}

	// make 100 deposits of $1
	// and 100 withdrawal of $1 concurrently.
	// run the program and check the result
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
