package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// https://newnews-api.herokuapp.com/
type News struct {
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Link      string `json:"link"`
		Points    int    `json:"points"`
		CreatedAt string `json:"created_at"`
	} `json:"items"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./app-name <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var news News

	err = json.Unmarshal(body, &news)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", news)
}
