package main

import "fmt"

func main() {
	str := "I'm a string"
	fmt.Println(str)
	fmt.Printf("%T\n", str)

	chars := []byte(str)
	fmt.Println(chars)
	fmt.Printf("%T\n", chars)
}