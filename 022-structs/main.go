package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type employee struct {
	person
	jobTitle string
}

func main() {
	john := person{
		first: "John",
		last:  "Doe",
		age:   0,
	}
	bob := person{
		first: "Bob",
		last:  "Duh",
		age:   40,
	}

	fmt.Println(john)
	fmt.Println(bob)
	fmt.Println()

	// Embedded structs
	employeeBob := employee{
		person:   bob,
		jobTitle: "Software Developer",
	}
	employeeAlice := employee{
		person: person{
			first: "Alice",
			last:  "Duh",
			age:   30,
		},
		jobTitle: "Software Developer",
	}
	fmt.Println(employeeBob)
	fmt.Println(employeeAlice)
	fmt.Println()
}
