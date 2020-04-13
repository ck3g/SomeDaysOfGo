package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	age   int
}

func main() {
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

	rawJSON, err := json.Marshal(people) // Does't marhall lowercase fields (because they are not exported)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n\n", string(rawJSON)) // Doesn't contain age (because age is not exported)

}
