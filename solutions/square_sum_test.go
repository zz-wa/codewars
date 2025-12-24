package solutions

import "testing"

func TestSquareSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2}, 5},
		{[]int{0, 3, 4, 5}, 50},
		{[]int{}, 0},
		{[]int{-1, -2}, 5},
		{[]int{1, 2, 2}, 9},
	}

	for _, test := range tests {
		result := SquareSum(test.input)
		if result != test.expected {
			t.Errorf("SquareSum(%v) = %d; want %d", test.input, result, test.expected)
		}
	}
}
