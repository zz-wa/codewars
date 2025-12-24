package solutions

import "testing"

func TestRomanEncoder(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1990, "MCMXC"},
		{2008, "MMVIII"},
		{1666, "MDCLXVI"},
		{3999, "MMMCMXCIX"},
	}

	for _, test := range tests {
		result := Solution(test.input)
		if result != test.expected {
			t.Errorf("Solution(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}
