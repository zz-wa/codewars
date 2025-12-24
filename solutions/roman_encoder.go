package solutions

import "strings"

type romanMapping struct {
	Value  int
	Symbol string
}

func Solution(number int) string {
	mappings := []romanMapping{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	for _, m := range mappings {
		for number >= m.Value {
			result.WriteString(m.Symbol)
			number -= m.Value
		}
	}

	return result.String()
}
