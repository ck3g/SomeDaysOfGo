package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("WebDev/005-predefined-functions-in-templates/*.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
