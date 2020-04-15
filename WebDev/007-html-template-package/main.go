package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("WebDev/007-html-template-package/tpl.gohtml"))
}

func main() {
	home := page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
