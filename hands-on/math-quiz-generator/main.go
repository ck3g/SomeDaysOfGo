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
		q := hardAdditionQuestion(100)
		fmt.Println(q)
	}
}

func easyAdditionQuestion() string {
	first := rand.Intn(9)
	second := rand.Intn(9)
	return fmt.Sprintf("%d + %d = __", first, second)
}

func mediumAdditionQuestion(maxSum int) string {
	sum := randomInRange(40, maxSum)
	first := randomInRange(10, sum-1)
	second := sum - first

	return fmt.Sprintf("%d + %d = __", first, second)
}

func hardAdditionQuestion(maxSum int) string {
	sum := randomInRange(25, maxSum)
	first := randomInRange(10, sum-1)
	second := sum - first

	return fmt.Sprintf("%d + %d = __", first, second)
}

func easySubtractionQuestion() string {
	first := randomInRange(10, 20)
	second := randomInRange(0, first)

	return fmt.Sprintf("%d - %d = __", first, second)
}

// mediumSubtractionQuestion always have a 10..20 minuend, and a 1..9 difference
func mediumSubtractionQuestion() string {
	diff := randomInRange(1, 9)
	minuend := randomInRange(10, 20)
	subtrahend := minuend - diff

	return fmt.Sprintf("%d - %d = __", minuend, subtrahend)
}

func easyMultiplicationQuestion() string {
	first := randomInRange(2, 10)
	second := randomInRange(2, 10)

	return fmt.Sprintf("%d · %d = __", first, second)
}

// randomInRange generates a random number in min...max range
func randomInRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
