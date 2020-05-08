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

	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	var cow = `         \  ^__^
	  \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
	      `

	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println(cow)
	fmt.Println()
}
