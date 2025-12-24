package solutions

func CountSheep(numbers []bool) int {
	count := 0
	for _, present := range numbers {
		if present {
			count++
		}
	}
	return count
}
