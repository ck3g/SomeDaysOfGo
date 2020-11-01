package main

import "fmt"

type person struct {
	first string
}

type cyclist struct {
	person
	bike string
}

func (p person) speak() {
	fmt.Printf("Hello, my name is %s.\n", p.first)
}

func (c *cyclist) speak() {
	fmt.Printf("Hello, my name is %s. I ride %s bike.\n", c.first, c.bike)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "John",
	}
	c := cyclist{
		person: person{
			first: "Bob",
		},
		bike: "Rondo Ruut AL 1",
	}

	saySomething(p)
	saySomething(&c)
}
