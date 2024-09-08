package p3279

import "sort"

func maxPossibleScore(start []int, d int) int {
	sort.Ints(start)
	n := len(start)
	// 二分吗？
	check := func(x int) bool {
		prev := start[0]

		for i := 1; i < n; i++ {
			cur := max(prev+x, start[i])
			if cur <= start[i]+d {
				prev = cur
			} else {
				return false
			}
		}
		return true
	}

	l, r := 0, start[n-1]+d+10
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return r - 1
}
