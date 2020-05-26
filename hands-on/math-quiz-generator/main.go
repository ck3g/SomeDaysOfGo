package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 60
	for i := 0; i < count; i++ {
		q := mediumAdditionQuestion()
		fmt.Println(q)
	}
}

func easyAdditionQuestion() string {
	first := rand.Intn(9)
	second := rand.Intn(9)
	return fmt.Sprintf("%d + %d = __", first, second)
}

func mediumAdditionQuestion() string {
	first := randomInRange(7, 20)
	min := 2

	if first < 10 {
		min = 8
	}

	second := randomInRange(min, 10)

	if rand.Intn(2) == 1 {
		first, second = second, first
	}

	return fmt.Sprintf("%d + %d = __", first, second)
}

// randomInRange generates a random number in min...max range
func randomInRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
