package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	example1()
}

func example1() {
	fmt.Printf("=== Example #1 =============\n\n")
	fmt.Println("Marshaling a JSON with unexported value (age)")

	type person struct {
		First string
		Last  string
		age   int
	}

	alice := person{
		First: "Alice",
		Last:  "White",
		age:   29,
	}
	bob := person{
		First: "Bob",
		Last:  "Blank",
		age:   36,
	}
	people := []person{alice, bob}
	fmt.Printf("%+v\n\n", people)

	jsonString, err := json.Marshal(people) // Does't marhall lowercase fields (because they are not exported)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n\n", string(jsonString)) // Doesn't contain age (because age is not exported)
}
