package solutions

import "sort"

func Gimme(array [3]int) int {
	// Create a slice from the array to sort it while keeping the original
	sorted := make([]int, 3)
	copy(sorted, array[:])
	sort.Ints(sorted)

	// The middle element in value is now at index 1 of the sorted slice
	middleValue := sorted[1]

	// Find the index of this value in the original array
	for i, v := range array {
		if v == middleValue {
			return i
		}
	}
	return -1 // Should not happen given constraints
}
