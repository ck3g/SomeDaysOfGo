package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println()

	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var count int

	return func() int {
		count++
		return count
	}
}
