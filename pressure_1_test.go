package kata

import "testing"

func TestDoubleInteger(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 4},
		{-1, -2},
		{0, 0},
		{100, 200},
	}

	for _, test := range tests {
		result := DoubleInteger(test.input)
		if result != test.expected {
			t.Errorf("DoubleInteger(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
