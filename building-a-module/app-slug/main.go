package main

import (
	"fmt"
	"log"

	"github.com/ck3g/SomeDaysOfGo/building-a-module/toolkit"
)

func main() {
	toSlug := "Now!!! is the time 123"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(slugified)
}
