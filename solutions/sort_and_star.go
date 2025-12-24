package solutions

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	s := arr[0]
	// Split the string into characters and join with "***"
	// However, strings.Split(s, "") gives empty strings at ends if not careful,
	// but for "abc", it gives ["a", "b", "c"].
	// Let's verify standard behavior or just manual loop.
	// strings.Split("abc", "") -> ["a", "b", "c"]
	
	chars := strings.Split(s, "")
	return strings.Join(chars, "***")
}
