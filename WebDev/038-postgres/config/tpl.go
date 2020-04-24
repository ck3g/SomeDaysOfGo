package config

import "html/template"

// TPL provides an access to HTML templates
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
