package main

import (
	"fmt"
	"log"
	"os"
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

// init function runs once when the program starts
// Can hold some initialization
func init() {
	tpl = template.Must(template.ParseGlob("WebDev/003-passing-data-into-templates/*.gohtml"))
}

func main() {
	fmt.Println("\n\n### Simple value")
	err := tpl.ExecuteTemplate(os.Stdout, "simple_data.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n### Using variable")

	err = tpl.ExecuteTemplate(os.Stdout, "using_variable.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n### Passing a slice")

	names := []string{"Alice", "Bob", "John", "Walter"}

	err = tpl.ExecuteTemplate(os.Stdout, "slices.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n### Passing a map")

	ages := map[string]int{
		"Alice":  29,
		"Bob":    35,
		"John":   34,
		"Walter": 52,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "maps.gohtml", ages)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n### Passing a struct")

	alice := person{
		Name: "Alice",
		Age:  29,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", alice)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n### Passing a slice of structs")

	people := []person{
		{"Alice", 29},
		{"Bob", 35},
		{"Walter", 52},
	}

	err = tpl.ExecuteTemplate(os.Stdout, "slice_struct.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n### Passing more complex data")

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

	err = tpl.ExecuteTemplate(os.Stdout, "slice_struct_slice.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
