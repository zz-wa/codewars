package solutions

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{2, "Even"},
		{0, "Even"},
		{7, "Odd"},
		{1, "Odd"},
		{-1, "Odd"}, // % in Go returns negative for negative operands (e.g. -1 % 2 = -1)
		             // Wait, -1 % 2 == -1. 
		             // My implementation check number%2 == 0. 
		             // -1 % 2 is -1, which is != 0, so it returns "Odd". Correct.
		             // -2 % 2 is 0. Returns "Even". Correct.
		{-2, "Even"},
	}

	for _, test := range tests {
		result := EvenOrOdd(test.input)
		if result != test.expected {
			t.Errorf("EvenOrOdd(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}
