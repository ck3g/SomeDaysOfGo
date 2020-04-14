package main

import "fmt"

func main() {
	fmt.Println("1 + 2 + 3 =", Sum(1, 2, 3))
}

// Sum calculates a sum of the numbers
func Sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}
