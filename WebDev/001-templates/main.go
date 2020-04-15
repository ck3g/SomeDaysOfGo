package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "John Doe"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello World!</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`

	// Print into a standard output
	fmt.Println(tpl)

	f, err := os.Create("WebDev/001-templates/index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer f.Close()

	io.Copy(f, strings.NewReader(tpl))
}
