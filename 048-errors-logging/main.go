package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error happened", err)
		fmt.Println()

		log.Println("Error happened", err) // By default prints in the standard output
	}

	f, err := os.Create("048-errors-logging/log")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Error happened", err) // Now prints into log, because of `log.SetOutput(f)`
		// log.Fatalln("Error happened (and program exited)", err)
		// log.Panicln("Error happened", err)
		panic(err)
	}
	defer f2.Close()

	fmt.Println()
	fmt.Println("check the log.txt file in the directory")
	fmt.Println()
}
