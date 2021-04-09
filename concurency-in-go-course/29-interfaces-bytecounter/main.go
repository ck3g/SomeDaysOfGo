package main

import "fmt"

type ByteCounter int

// Implement Write method for ByteCounter to count the number of bytes written
func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
