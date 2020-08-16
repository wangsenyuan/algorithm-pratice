package p1551

func minOperations(n int) int {
	// 2 * i + 1
	// 2 * (n - 1 - i) + 1 = 2 * n - 2 * i - 1 = 2 * n - (2 * i + 1)
	// i + (n - 1 - i) = 2 * n

	var res int

	for i := 0; i < n; i++ {
		x := 2*i + 1
		if x >= n {
			break
		}
		res += n - x
	}
	return res
}
