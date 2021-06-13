package p1894

func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	var sum int

	for i := 0; i < n; i++ {
		sum += chalk[i]
		if sum > k {
			return i
		}
	}
	k %= sum
	sum = 0
	for i := 0; i < n; i++ {
		sum += chalk[i]
		if sum > k {
			return i
		}
	}
	return 0
}
