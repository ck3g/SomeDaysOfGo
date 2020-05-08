package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
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

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxwidth)

	for _, message := range messages {
		fmt.Println(message)
	}

	fmt.Println(cow)
	fmt.Println()
}

// tabsToSpaces converts all tabs found in the strings
// found in the `lines` slice to 4 spaces, to prevent misalignments in
// counting the runes
func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

// calculateMaxWidth given a slice of strings return the length of the
// string with max length
func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return w
}

// normalizeStringsLength takes a slice of strings and appends
// to each one a numbef of spaces needed to have them all the same number
// of runes
func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}

	return ret
}
