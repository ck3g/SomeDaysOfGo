package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "Practice makes perfect"
	encoded := encode(msg)
	fmt.Println("ENCODED MESSAGE", encoded)

	str, err := decode(encoded)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("DECODED MESSAGE", str)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encoded string) (string, error) {
	bs, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("Couldn't decode string %w", err)
	}

	return string(bs), nil
}
