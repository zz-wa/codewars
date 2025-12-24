package solutions

import "testing"

func TestGetCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abracadabra", 5},
		{"pear tree", 4},
		{"o a kak ushakov lil vo kashu kakao", 13},
		{"my pyx", 0},
		{"", 0},
	}

	for _, test := range tests {
		result := GetCount(test.input)
		if result != test.expected {
			t.Errorf("GetCount(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
