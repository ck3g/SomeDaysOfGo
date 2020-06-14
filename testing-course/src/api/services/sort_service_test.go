package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// That is an integration test. Even though technically there is no difference with unit test
func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	Sort(elements)

	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.EqualValues(t, want, elements)
}
