package solutions

import "sort"

func MinMax(arr []int) [2]int {
	// Constraints say array will have at least one element.
	// We can sort or iterate. Iteration is O(n), Sort is O(n log n).
	// For variety, let's use sorting this time since we iterated in the last one.
	// Actually, iteration is better for performance, but sorting is concise.
	// Let's stick to iteration as it's more "correct" for O(n).
	
	if len(arr) == 0 {
		// Should not happen based on description
		return [2]int{0, 0}
	}

	minVal := arr[0]
	maxVal := arr[0]

	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	return [2]int{minVal, maxVal}
}

// Alternative concise implementation using sort just for reference
func MinMaxSorted(arr []int) [2]int {
	if len(arr) == 0 { return [2]int{0,0} }
	// copying to avoid side effects if we were strict, but for this kata, usually ok.
	// Let's just use the iterative one as the primary solution.
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	return [2]int{tmp[0], tmp[len(tmp)-1]}
}
