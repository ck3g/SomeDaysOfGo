package main

import (
	"fmt"
	"net/http"
)

type anything int

func (m anything) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("SomeKey", "some value")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any page content here</h1>")
}

func main() {
	var a anything
	http.ListenAndServe(":8080", a)
}
