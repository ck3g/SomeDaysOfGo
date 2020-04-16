package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Home")
}

func about(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "About")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080", nil)
}
