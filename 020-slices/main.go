package main

import "fmt"

func main() {
	// Slice allows to group values of the same type
	x := []int{2, 4, 8, 16, 32, 64, 128, 256}
	fmt.Println(x)

	fmt.Printf("len(x) = %v\n\n", len(x))

	// Looping through slice
	for index, value := range x {
		fmt.Printf("x has value %v at index %v\n", value, index)
	}
	fmt.Println()

	// Slicing a slice
	fmt.Println(x[1:])  // [4 8 16 32 64 128 256]
	fmt.Println(x[:4])  // [2 4 8 16]
	fmt.Println(x[2:6]) // [8 16 32 64]
	fmt.Println()

	// Append to a slice
	x = append(x, 512, 1024)
	fmt.Println(x)
	fmt.Println()

	y := []int{1, 2, 3}
	z := []int{10, 20, 30}
	y = append(y, z...)
	fmt.Println(y)
	fmt.Println()

	// Delete from a slice
	y = append(y[:2], y[4:]...) // use 0-2 slice and merge it with 4-end slice
	fmt.Println(y)              // [1 2 20 30]
	fmt.Println()

	// make slice
	a := make([]int, 10, 12)
	fmt.Println("'a' is", a)
	fmt.Printf("length of 'a' is %v\n", len(a))
	fmt.Printf("capacity of 'a' is %v\n", cap(a))
	fmt.Println()

	// a[10] = 1 // Can't do: index out of range [10] with length 10
	a = append(a, 1, 1)
	fmt.Println("'a' is", a)
	fmt.Printf("length of 'a' is %v\n", len(a))
	fmt.Printf("capacity of 'a' is %v\n", cap(a))
	fmt.Println()

	a = append(a, 2) // If length becames more than a capacity, then capacity a slice doubles in size
	fmt.Println("'a' is", a)
	fmt.Printf("length of 'a' is %v\n", len(a))
	fmt.Printf("capacity of 'a' is %v\n", cap(a))
	fmt.Println()

	// Multi-dimentional slice
	// First name, last name, age, height
	johnDoe := []string{"John", "Doe", "N/A", "N/A"}
	bobUndoe := []string{"Bob", "Undoe", "40", "180"}

	people := [][]string{johnDoe, bobUndoe}
	fmt.Println("people", people)
	fmt.Println()
}
