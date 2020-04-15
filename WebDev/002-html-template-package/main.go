package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	// Parse a template from a file
	tpl, err := template.ParseFiles("WebDev/002-html-template-package/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the template into a standart output
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the template into a file
	f, err := os.Create("WebDev/002-html-template-package/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer f.Close()

	err = tpl.Execute(f, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n------------------")

	// Once we have a template we can parse more files into it
	tpl, err = tpl.ParseFiles("WebDev/002-html-template-package/description.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Now we can execute a specific template
	err = tpl.ExecuteTemplate(os.Stdout, "description.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n------------------")

	// Here we parse a bunch of files using ParseGlob
	tpl, err = template.ParseGlob("WebDev/002-html-template-package/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()

	err = tpl.ExecuteTemplate(os.Stdout, "description.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\n------------------")

	// We can use teamplate.Must to skip error checking block
	tpl = template.Must(template.ParseGlob("WebDev/002-html-template-package/description.gohtml"))

	// Now we can execute a specific template
	err = tpl.ExecuteTemplate(os.Stdout, "description.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
