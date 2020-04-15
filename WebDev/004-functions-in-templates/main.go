package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

type company struct {
	Name       string
	HeadOffice string
}

var fm = template.FuncMap{
	"toUpper": strings.ToUpper,
	"first3":  first3,
}

func first3(str string) string {
	str = strings.TrimSpace(str)
	str = str[:3]
	return str
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("WebDev/004-functions-in-templates/*.gohtml"))
}

func main() {
	fmt.Println("\n\n### Passing more complex data")

	people := []person{
		{"Alice", 29},
		{"Bob", 35},
		{"Walter", 52},
	}

	companies := []company{
		{"Microsoft", "Redmond"},
		{"Apple", "Cupertino"},
		{"Google", "Mountain View"},
	}

	data := struct {
		People    []person
		Companies []company
	}{
		people,
		companies,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "function.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
