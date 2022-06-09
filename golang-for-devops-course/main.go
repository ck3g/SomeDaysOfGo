package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Printf("Hello world!\nArguments: %v\n", args)
}
