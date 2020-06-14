package sort

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// -- Initialization
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// -- Execution
	BubbleSort(elements)

	// -- Validation

	// Version from the course.
	// Checking first and last elements aren't enough, that doesn't guarantee that
	// the rest of the elements are sorted.
	// For example the initial slice (unsorted) will pass here.
	if elements[0] != 0 {
		t.Errorf("first element should be 0, got %v", elements[0])
	}
	if elements[len(elements)-1] != 9 {
		t.Errorf("last emenet should be 9, got %v", elements[len(elements)-1])
	}

	// Better version could be checking against whole slice
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := elements
	if !reflect.DeepEqual(got, want) {
		t.Errorf("elements are not sorted. got %v want %v", got, want)
	}

	// Even better using stretchr/testify
	assert.EqualValues(t, want, got)
}

func TestBubbleSort_WithTimeout(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	timeoutChan := make(chan bool, 1)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		t.Errorf("Bubble sort took more than 500ms")
		return
	}

	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.EqualValues(t, want, elements)
}

func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	Sort(elements)

	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.EqualValues(t, want, elements)
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
