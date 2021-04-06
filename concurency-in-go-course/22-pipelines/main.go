package main

import "fmt"

// TODO: build a pipeline
// generator() -> square() -> print

// generator - converts a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

// square - receives an inbound channel
// squares a number
// outputs an outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline'
	// receives the values from square stage
	// print each one, until channel is closed

	// one way
	// ch := generator(2, 3)
	// out := square(ch)
	// for n := range out {
	// 	fmt.Println(n)
	// }

	// another way
	for n := range square(generator(2, 3)) {
		fmt.Println(n)
	}

	// with double square
	for n := range square(square(generator(2, 3))) {
		fmt.Println(n)
	}
}
