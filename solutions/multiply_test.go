package solutions

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 6},
		{5, 0, 0},
		{-1, 5, -5},
		{10, 10, 100},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}
