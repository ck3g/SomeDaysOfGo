package main

import "fmt"

func main() {
	// init a map
	m := map[string]int{
		"Alice": 29,
		"Bob":   35,
		"John":  40,
	}

	// access keys
	fmt.Println(m)
	fmt.Println(m["Bob"])
	fmt.Println(m["Alice"])
	fmt.Println()

	// check if the key exists
	age, ok := m["Anonymous"]
	fmt.Println("Anonymous age", age, "exists?", ok)
	fmt.Println()

	// adding new values to a map
	m["Mike"] = 50
	fmt.Println(m)
	fmt.Println()

	// iterate over maps
	for name, age := range m {
		fmt.Printf("%v is %v\n", name, age)
	}
	fmt.Println()

	// Delete an item from a map
	delete(m, "Bob")
	delete(m, "John Doe") // Doesn't raise an error if the key does not exist
	fmt.Println(m)
	fmt.Println()
}
