package main

import "fmt"

type person struct {
	first string
	last  string
}

type employee struct {
	person
	jobTitle string
}

func (e employee) fullName() string {
	return e.first + " " + e.last
}

func main() {
	john := employee{
		person: person{
			first: "John",
			last:  "Doe",
		},
		jobTitle: "N/A",
	}
	bob := employee{
		person: person{
			first: "Bob",
			last:  "Duh",
		},
		jobTitle: "Software Developer",
	}

	fmt.Println(john)
	fmt.Println(bob)
	fmt.Println()

	fmt.Println(john.fullName())
	fmt.Println(bob.fullName())
	fmt.Println()
}
