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
	// curl -XGET -H "Content-type: application/json" -d '{"First":"Bob"}' localhost:8080/
	http.HandleFunc("/", doDecode)
	http.ListenAndServe(":8080", nil)
}

func doDecode(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}

	log.Println("Person", p1)
}
