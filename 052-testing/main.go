package main

import "fmt"

func main() {
	msg := Message("Here", "is", "the", "message")
	fmt.Println(msg)
}

// Message converts a list of string arguments into a string
func Message(m ...string) string {
	var message string
	for i, substr := range m {
		if i == 0 {
			message += substr
		} else {
			message += " " + substr
		}
	}

	return message
}
