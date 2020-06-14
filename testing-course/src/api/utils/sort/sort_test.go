package sort

import (
	"reflect"
	"testing"
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
}
