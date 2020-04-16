package main

import (
	"io"
	"net/http"
)

type anything int

func (m anything) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(res, "Home page")
	case "/about":
		io.WriteString(res, "About page")
	default:
		io.WriteString(res, "404: Page Not Found")
	}
}

func main() {
	var a anything
	http.ListenAndServe(":8080", a)
}
