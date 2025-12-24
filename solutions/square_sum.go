package solutions

func SquareSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n * n
	}
	return sum
}
