package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	parts := strings.Fields(in)
	if len(parts) == 0 {
		return ""
	}

	minVal, _ := strconv.Atoi(parts[0])
	maxVal := minVal

	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	return fmt.Sprintf("%d %d", maxVal, minVal)
}
