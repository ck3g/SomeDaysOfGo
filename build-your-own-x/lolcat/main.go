package main

import (
	"fmt"
	"math"
	"strings"

	"syreclabs.com/go/faker"
)

func main() {
	var phrases []string
	for i := 1; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	output := strings.Join(phrases[:], "; ")

	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
	fmt.Println()
}

func rgb(i int) (int, int, int) {
	f := 0.1
	r := int(math.Sin(f*float64(i)+0)*127 + 128)
	g := int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)

	return r, g, b
}
