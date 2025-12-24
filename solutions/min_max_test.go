package solutions

import "testing"

func TestMinMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected [2]int
	}{
		{[]int{1, 2, 3, 4, 5}, [2]int{1, 5}},
		{[]int{2334454, 5}, [2]int{5, 2334454}},
		{[]int{1}, [2]int{1, 1}},
		{[]int{5, 23, 4, 233, 10, -5}, [2]int{-5, 233}},
	}

	for _, test := range tests {
		result := MinMax(test.input)
		if result != test.expected {
			t.Errorf("MinMax(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
