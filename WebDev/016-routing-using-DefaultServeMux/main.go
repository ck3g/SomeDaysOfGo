package main

import (
	"io"
	"net/http"
)

type anything int

func (a anything) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Anything")
}

type anythingElse int

func (a anythingElse) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Anything else")
}

func main() {
	var a anything
	var e anythingElse

	http.Handle("/", a)
	http.Handle("/about", e)

	http.ListenAndServe(":8080", nil)
}
