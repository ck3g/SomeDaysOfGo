package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "secret123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Original password:", password)
	fmt.Println("  Hashed password:", string(hashedPassword))
	fmt.Println()

	fmt.Println("New login attempt...")
	newPassword := "wrong password"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(newPassword))
	if err != nil {
		fmt.Println("Incorrect password. Try again")
		fmt.Println()
	}

	fmt.Println("New login attempt...")
	newPassword = "secret123"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(newPassword))
	if err != nil {
		fmt.Println("Incorrect password. Try again")
		fmt.Println()
	} else {
		fmt.Println("Your password is correct")
	}
}
