package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	msg := "Practice makes perfect"

	password := "go-in-practice"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatalln("couldn't bcrypt password", err)
	}
	bs = bs[:16]

	rslt, err := enDencode(bs, msg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rslt))

	rslt2, err := enDencode(bs, string(rslt))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rslt2))
}

func enDencode(key []byte, input string) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("coulnd't newCipher %w", err)
	}

	// initialization vector
	iv := make([]byte, aes.BlockSize)

	// create a cipher
	s := cipher.NewCTR(b, iv)
	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter{
		S: s,
		W: buff,
	}
	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("couldn't sw.Write to StreamWriter %w", err)
	}

	return buff.Bytes(), nil
}
