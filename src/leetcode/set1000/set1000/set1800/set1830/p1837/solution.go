package p1837

func sumBase(n int, k int) int {
	// a0 * 1 + a1 * k + a2 * kk + a3 * kkk = n
	// a0 < k
	var sum int
	for n > 0 {
		r := n % k
		sum += r
		n -= r
		n /= k
	}
	return sum
}
