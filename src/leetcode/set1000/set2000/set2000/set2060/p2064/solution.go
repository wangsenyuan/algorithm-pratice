package p2064

func minimizedMaximum(n int, quantities []int) int {
	m := len(quantities)
	// whether can distribute at most x
	check := func(x int) bool {
		var cnt int

		for i := 0; i < m; i++ {
			cur := quantities[i]
			cnt += (cur + x - 1) / x
		}
		return cnt <= n
	}

	var l, r = 1, 200000
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
