package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := 60
	questions := easyAdditionQuestions(count)
	for _, q := range questions {
		fmt.Println(q)
	}
}

func easyAdditionQuestions(count int) []string {
	var questions []string
	for i := 0; i < count; i++ {
		questions = append(questions, easyAdditionQuestion())
	}
	return questions
}

func easyAdditionQuestion() string {
	first := rand.Intn(9)
	second := rand.Intn(9)
	return fmt.Sprintf("%d + %d = __", first, second)
}
