package main

import (
	"fmt"
	"net/http"
)

/*

Any type that implements the following interface can act as a handler

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

type myInt int

func (m myInt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want here")
}

func main() {
	var i myInt
	http.ListenAndServe(":8080", i)
}
