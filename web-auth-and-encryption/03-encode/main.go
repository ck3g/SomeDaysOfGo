package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/", doEncode)
	http.ListenAndServe(":8080", nil)
}

func doEncode(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Bob",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}
