package p2897

func differenceOfSums(n int, m int) int {
	sum := (1 + n) * n / 2
	var sum2 int

	for i := m; i <= n; i++ {
		if i%m == 0 {
			sum -= i
			sum2 += i
		}
	}

	return sum - sum2
}
