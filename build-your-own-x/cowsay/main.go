package main

import (
	"bufio"
	"flag"
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

	var figure string

	flag.StringVar(&figure, "f", "cow", "the figure name. Valid values are `cow` and `stegosaurus`")
	flag.Parse()

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)

	fmt.Println(balloon)
	printFigure(figure)
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

// buildBalloon takes a slice of strings of max width maxwidth
// prepends/appends margins of first and last line, and at start/end of each line
// and returns a string with the contents of the balloon
func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

// printFigure given a figure name prints it.
// Currently accepts `cow` and `stegosaurus`.
func printFigure(name string) {
	var cow = `       \  ^__^
	\ (oo)\_______
	  (__)\       )\/\
	      ||----w |
	      ||     ||
	    `

	var stegosaurus = `           \                      .       .
	    \                    / ` + "`" + `.   .' "
	     \           .---.  <    > <    >  .---.
	      \          |    \  \ - ~ ~ - /  /    |
	    _____           ..-~             ~-..-~
	   |     |   \~~~\\.'                    ` + "`" + `./~~~/
	  ---------   \__/                         \__/
	 .'  O    \     /               /       \  "
	(_____,    ` + "`" + `._.'               |         }  \/~~~/
	 ` + "`" + `----.          /       }     |        /    \__/
	       ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
		   ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
			|     /        |     /     ~-.     ` + "`" + `-. _  _  _
			|_____|        |_____|         ~ - . _ _ _ _ _>
  
	  `

	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	default:
		fmt.Println("Unknown figure")
	}
}
