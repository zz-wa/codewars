package kata

import (
	"testing"
)

func TestGimme(t *testing.T) {
	tests := []struct {
		input    [3]int
		expected int
	}{
		{[3]int{2, 3, 1}, 0},
		{[3]int{5, 10, 14}, 1},
		{[3]int{1, 3, 4}, 1},
		{[3]int{15, 10, 14}, 2},
		{[3]int{-4, -23, 4}, 0},
		{[3]int{-15, -10, 14}, 1},
	}

	for _, test := range tests {
		result := Gimme(test.input)
		if result != test.expected {
			t.Errorf("Gimme(%v) = %d; want %d", test.input, result, test.expected)
		}
	}
}
