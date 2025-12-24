package solutions

import "testing"

func TestCountSheep(t *testing.T) {
	tests := []struct {
		input    []bool
		expected int
	}{
		{[]bool{true, true, true, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, false, true, true}, 17},
		{[]bool{false}, 0},
		{[]bool{true}, 1},
		{[]bool{}, 0},
	}

	for _, test := range tests {
		result := CountSheep(test.input)
		if result != test.expected {
			t.Errorf("CountSheep(%v) = %d; want %d", test.input, result, test.expected)
		}
	}
}
