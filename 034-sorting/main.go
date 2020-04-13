package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 7, 1, 5, 8, 2, 6}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println()

	s := []string{"Walter", "Alice", "John", "Bob"}
	sort.Strings(s)
	fmt.Println(s)
	fmt.Println()
}
