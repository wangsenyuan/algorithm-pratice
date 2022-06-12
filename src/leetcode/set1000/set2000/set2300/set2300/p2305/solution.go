package p2305

func distributeCookies(cookies []int, k int) int {
	var l, r int

	for _, cur := range cookies {
		r += cur
	}

	for l < r {
		mid := (l + r) / 2
		if check(cookies, k, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func check(cookies []int, k int, expect int) bool {
	n := len(cookies)
	var loop func(i int) bool

	sum := make([]int, k)

	loop = func(i int) bool {
		if i == n {
			return true
		}
		// distribute current cookie to child
		for j := 0; j < k; j++ {
			if sum[j]+cookies[i] <= expect {
				sum[j] += cookies[i]
				if loop(i + 1) {
					return true
				}
				sum[j] -= cookies[i]
			}
		}
		return false
	}

	res := loop(0)

	return res
}
