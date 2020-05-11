package main

import (
	"flag"
	"fmt"
)

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}

// scan given a path crawls it and its subfolders
// searching for Git repositories
func scan(folder string) {
	fmt.Printf("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

// stats generates a nice graph of your Git contributions
func stats(email string) {
	print("stats")
}

func recursiveScanFolder(folder string) []string {
	var folders []string
	return folders
}

func getDotFilePath() string {
	return ""
}

func addNewSliceElementsToFile(filePath string, repositories []string) {

}
