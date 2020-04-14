package somedays

import "fmt"

// Also generates documentation for the Message function
func ExampleMessage() {
	fmt.Println(Message("Here", "is", "the", "message"))
	// Output:
	// Here is the message
}
