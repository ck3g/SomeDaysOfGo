package services

import (
	"reflect"
	"testing"
)

// That is an integration test. Even though technically there is no difference with unit test
func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	Sort(elements)

	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	got := elements
	if !reflect.DeepEqual(got, want) {
		t.Errorf("elements are not sorted. got %v want %v", got, want)
	}
}
