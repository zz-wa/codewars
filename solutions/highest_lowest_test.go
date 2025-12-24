package solutions

import "testing"

func TestHighAndLow(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 2 3 4 5", "5 1"},
		{"1 2 -3 4 5", "5 -3"},
		{"1 9 3 4 -5", "9 -5"},
		{"8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
		{"1 2 3", "3 1"},
	}

	for _, test := range tests {
		result := HighAndLow(test.input)
		if result != test.expected {
			t.Errorf("HighAndLow(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
