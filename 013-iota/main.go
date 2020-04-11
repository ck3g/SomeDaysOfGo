package main

import "fmt"

const (
	a = iota
	b = iota
	c
	d
)

const (
	e = iota
	f
)

func main() {
	fmt.Printf("a = %v (%T)\n", a, a)
	fmt.Printf("b = %v (%T)\n", b, b)
	fmt.Printf("c = %v (%T)\n", c, c)
	fmt.Printf("d = %v (%T)\n", d, d)
	fmt.Printf("e = %v (%T)\n", e, e)
	fmt.Printf("f = %v (%T)\n", f, f)
}