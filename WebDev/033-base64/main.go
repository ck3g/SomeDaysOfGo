package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Hey Ho Let's Go!"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
	fmt.Println()

	s64 = base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)
	fmt.Println()

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("Can't go!")
	}
	fmt.Println(string(bs))
	fmt.Println()
}
