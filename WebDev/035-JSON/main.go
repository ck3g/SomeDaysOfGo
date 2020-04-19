/*
The most important thing to understand is that you can marshal *OR* encode Go code to JSON.
Regardless of whether or not you use “marshal” or “encode”, your Go data structures will be turned into JSON.
So what’s the difference? Marshal is for turning Go data structures into JSON and then assigning the JSON to a variable.
Encode is used to turn Go data structures into JSON and then send it over the wire.
Both" marshal" and "encode" have their counterparts: "unmarshal" and "decode".
You can learn about Go & JSON at https://godoc.org/encoding/json - Package json implements encoding and decoding of
JSON as defined in RFC 4627.
The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.
You can also read about Go & JSON at this Go official blogpost: https://blog.golang.org/json-and-go

*/

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
	Items []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", jsonMarshal)
	http.HandleFunc("/encode", jsonEncode)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `
		<h2>Here is the things to try</h2>
		<div><a href="/marshal">Marshal</a></div>
		<div><a href="/encode">Encode</a></div>
	`)
}

func jsonMarshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	johnDoe := person{
		First: "John",
		Last:  "Doe",
		Items: []string{"key", "gum", "sun glasses"},
	}
	j, err := json.Marshal(johnDoe)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func jsonEncode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	johnDoe := person{
		First: "John",
		Last:  "Doe",
		Items: []string{"key", "gum", "sun glasses"},
	}
	err := json.NewEncoder(w).Encode(johnDoe)
	if err != nil {
		log.Println(err)
	}
}
