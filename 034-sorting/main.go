package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

// ByAge implements sort.Interface for [] person
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	// Sorting integers
	a := []int{3, 7, 1, 5, 8, 2, 6}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println()

	// Sorting a slice of strings
	s := []string{"Walter", "Alice", "John", "Bob"}
	sort.Strings(s)
	fmt.Println(s)
	fmt.Println()

	// Sorting by custom field
	people := []person{
		{"Walter White", 52},
		{"Alice Pink", 29},
		{"John Doe", 25},
	}
	fmt.Printf("%+v\n", people)
	sort.Sort(ByAge(people))
	fmt.Printf("%+v\n", people)
}
