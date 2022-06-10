package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./app-name <argument>")
		os.Exit(1)
	}

	fmt.Printf("Hello world!\nArguments: %v\n", args[1:])
}
