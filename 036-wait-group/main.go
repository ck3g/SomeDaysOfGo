package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Add(1)
	go foo() // 'go' runs a concurrent function (goroutine)
	bar()

	fmt.Println("\nGoroutines\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Wait() // Wait for goroutine to finish

	fmt.Println("\nGoroutines\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Print("*")
	}
	wg.Done() // Notify that the goroutine is finished
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Print(".")
	}
}
