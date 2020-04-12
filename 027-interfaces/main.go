package main

import "fmt"

type person struct {
	first string
	last  string
}

func (e person) fullName() string {
	return e.first + " " + e.last
}

type employee struct {
	person
	jobTitle string
}

func (e employee) fullName() string {
	return e.first + " " + e.last + " [" + e.jobTitle + "]"
}

type human interface {
	fullName() string
}

func greet(h human) {
	fmt.Println("Hello ", h.fullName())
}

func main() {
	alice := person{
		first: "Alice",
		last:  "White",
	}
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

	fmt.Println(alice)
	fmt.Println(john)
	fmt.Println(bob)
	fmt.Println()

	alice.fullName()
	john.fullName()
	bob.fullName()
	fmt.Println()

	greet(alice)
	greet(john)
	greet(bob)
}
