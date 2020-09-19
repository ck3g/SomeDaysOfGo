package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	rawJSON := `[{"First": "Bob"}, {"First": "Alice"}]`
	xp := []person{}
	err := json.Unmarshal([]byte(rawJSON), &xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("From JSON to Go stuct: %+v\n", xp)
}
