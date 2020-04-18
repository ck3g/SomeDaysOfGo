package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	code := getCode("test@example.com")
	fmt.Println(code)
	code = getCode("test@exampl.com")
	fmt.Println(code)
}

func getCode(str string) string {
	hash := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
