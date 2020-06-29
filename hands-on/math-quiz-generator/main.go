package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 30
	for i := 0; i < count; i++ {
		q := extraHardSubtractionQuestion()
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

func hardSubtractionQuestion() string {
	diff := randomInRange(1, 50)
	minuend := randomInRange(diff, 99)
	subtrahend := minuend - diff

	return fmt.Sprintf("%d - %d = __", minuend, subtrahend)
}

func extraHardSubtractionQuestion() string {
	subtrahendOnes := randomInRange(2, 9)
	minuendOnes := randomInRange(1, subtrahendOnes-1)
	minuendDozen := randomInRange(2, 9)
	subtrahendDozen := randomInRange(1, minuendDozen-1)

	minuend := minuendDozen*10 + minuendOnes
	subtrahend := subtrahendDozen*10 + subtrahendOnes

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
