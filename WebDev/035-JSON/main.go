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
	"fmt"
	"io"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
	Items []string
}

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

type city struct {
	Latitude, Longitude float64
	City                string
}

type cities []city

type taggedCities []struct {
	Postal    string  `json:"postal"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Zip       string  `json:"zip"`
	Country   string  `json:"country"`
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", jsonMarshal)
	http.HandleFunc("/encode", jsonEncode)
	http.HandleFunc("/unmarshal", jsonUnmarshal)
	http.HandleFunc("/unmarshal-array", jsonUnmarshalArray)
	http.HandleFunc("/unmarshal-array-tagged", jsonUnmarshalArrayTagged)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `
		<h2>Here is the things to try</h2>
		<div><a href="/marshal">Marshal</a></div>
		<div><a href="/encode">Encode</a></div>
		<div><a href="/unmarshal">Unmarshal</a></div>
		<div><a href="/unmarshal-array">Unmarshal an Array</a></div>
		<div><a href="/unmarshal-array-tagged">Unmarshal an Array Tagged</a></div>
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

func jsonUnmarshal(w http.ResponseWriter, req *http.Request) {
	var data img
	received := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(received), &data)
	if err != nil {
		log.Fatalln("error unmarshaling", err)
	}

	fmt.Fprintf(w, "%+v\n", data)

	for i, v := range data.IDs {
		fmt.Fprintln(w, i, v)
	}

	fmt.Fprintln(w, data.Thumbnail.URL)
}

func jsonUnmarshalArray(w http.ResponseWriter, req *http.Request) {
	var data cities
	received := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(received), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, received)
	fmt.Fprintf(w, "%+v\n", data)
}

func jsonUnmarshalArrayTagged(w http.ResponseWriter, req *http.Request) {
	var data taggedCities
	received := `[{"postal":"zip","latitude":37.7668,"longitude":-122.3959,"address":"","city":"SAN FRANCISCO","state":"CA","zip":"94107","country":"US"},{"postal":"zip","latitude":37.371991,"longitude":-122.02602,"address":"","city":"SUNNYVALE","state":"CA","zip":"94085","country":"US"}]`

	err := json.Unmarshal([]byte(received), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, received)
	fmt.Fprintf(w, "%+v\n", data)
}
