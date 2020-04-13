package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	quit <- 0
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("From the even channel:", v)
		case v := <-odd:
			fmt.Println("From the odd channel:", v)
		case v := <-quit:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}
