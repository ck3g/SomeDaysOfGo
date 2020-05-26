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
		q := mediumAdditionQuestion(20)
		fmt.Println(q)
	}
}

func easyAdditionQuestion() string {
	first := rand.Intn(9)
	second := rand.Intn(9)
	return fmt.Sprintf("%d + %d = __", first, second)
}

func mediumAdditionQuestion(maxSum int) string {
	sum := randomInRange(10, maxSum)
	first := randomInRange(2, sum-1)
	second := sum - first

	return fmt.Sprintf("%d + %d = __", first, second)
}

// randomInRange generates a random number in min...max range
func randomInRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
