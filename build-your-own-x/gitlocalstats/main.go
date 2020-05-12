package main

import (
	"flag"
)

var repositories []string

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		// return // Not storing to the file, so process every time
	}

	stats(email)
}
