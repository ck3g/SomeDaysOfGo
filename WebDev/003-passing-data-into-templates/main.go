package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// init function runs once when the program starts
// Can hold some initialization
func init() {
	tpl = template.Must(template.ParseGlob("WebDev/003-passing-data-into-templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "simple_data.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n-------")

	err = tpl.ExecuteTemplate(os.Stdout, "using_variable.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
