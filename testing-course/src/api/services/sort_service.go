package services

import "github.com/ck3g/SomeDaysOfGo/testing-course/src/api/utils/sort"

// Sort provides a sort function
func Sort(elements []int) {
	sort.BubbleSort(elements)
}
