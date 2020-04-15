package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("WebDev/006-nesting-templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "layout.gohtml", "dynamic content")
	if err != nil {
		log.Fatalln(err)
	}
}
