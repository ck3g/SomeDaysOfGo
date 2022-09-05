package main

import (
	"fmt"

	"github.com/ck3g/SomeDaysOfGo/building-a-module/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("Random string:", s)
}
