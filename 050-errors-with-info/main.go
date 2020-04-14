package main

import (
	"errors"
	"fmt"
	"log"
)

// ErrMath contains an error example
var ErrMath = errors.New("math: square root of negatve number")

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("math: square root of negatve number")
		// return 0, ErrMath
		return 0, fmt.Errorf("math: square root of negatve number: %v", f)
	}

	return 503, nil
}
