package p2829

func minimumSum(n int, k int) int {
	// (a + b ) % k == 0
	// 则不能同时选择 a, b
	use := make(map[int]bool)
	var sum int
	cur := 1
	for i := 1; i <= n; i++ {
		for k-cur >= 0 && use[k-cur] {
			cur++
		}
		use[cur] = true
		sum += cur
		cur++
	}
	return sum
}
