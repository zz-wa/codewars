package solutions

import "testing"

func TestTwoSort(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}, "b***i***t***c***o***i***n"},
		{[]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}, "a***r***e"},
		{[]string{"lets", "talk", "about", "javascript", "the", "best", "language"}, "a***b***o***u***t"},
		{[]string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}, "c***o***d***e"},
		{[]string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}, "L***e***t***s"},
	}

	for _, test := range tests {
		// Make a copy strictly to follow function signature of passing a slice,
		// though sorting modifies it in place, it's fine for the test nature.
		// However, to be safe against side-effects if we re-used slices (we don't here),
		// we just pass it directly.
		inputCopy := make([]string, len(test.input))
		copy(inputCopy, test.input)

		result := TwoSort(inputCopy)
		if result != test.expected {
			t.Errorf("TwoSort(%v) = %q; want %q", test.input, result, test.expected)
		}
	}
}
