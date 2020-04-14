package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println()

	// ------

	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Favorite food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)
	fmt.Println()

	// -------------

	f, err := os.Create("047-error-checking/names.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := strings.NewReader("John Doe")
	io.Copy(f, r)
	fmt.Println()

	// ---------

	f, err = os.Open("047-error-checking/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	fmt.Println()
}
