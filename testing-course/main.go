package main

import (
	"fmt"

	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/providers/locations_provider"
)

func main() {
	c, err := locations_provider.GetCountry("DEU")
	if err != nil {
		fmt.Println(err.Message)
		return
	}

	fmt.Printf("Country: %+v", c)
}
