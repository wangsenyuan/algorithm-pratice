package p3449

import "slices"

func maxScore(points []int, m int) int64 {
	n := len(points)
	if m < n {
		return 0
	}

	check := func(exp int) bool {
		var prev int
		cnt := m
		for i, p := range points {
			k := (exp-1)/p + 1 - prev
			if i == n-1 && k <= 0 {
				break
			}
			k = max(k, 1)
			cnt -= 2*k - 1
			if cnt < 0 {
				return false
			}
			prev = k - 1
		}

		return true
	}
	l, r := 1, 1+(m+1)/2*slices.Min(points)
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return int64(r - 1)
}
