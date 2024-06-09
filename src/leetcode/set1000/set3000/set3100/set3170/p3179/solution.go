package p3179

func numberOfChild(n int, k int) int {
	r := k % (2 * (n - 1))

	if r >= n-1 {
		r -= n - 1
		r = n - 1 - r
	}

	return r
}
