package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat() // Stat() returns FileInfo https://godoc.org/os#FileInfo

	if info.Mode()&os.ModeCharDevice != 0 {
		// https://godoc.org/os#FileMode
		// os.ModeCharDevice // c: Unix character device, when ModeDevice is set
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println(`Usage: echo "<TEXT>" | cowsay`)
		return
	}

	reader := bufio.NewReader(os.Stdin) // https://godoc.org/bufio#NewReader
	var output []rune

	for {
		input, _, err := reader.ReadRune() // https://godoc.org/bufio#Reader.ReadRune
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
}
