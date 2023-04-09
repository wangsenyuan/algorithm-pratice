package p2616

import (
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)

	n := len(nums)

	check := func(d int) bool {
		var cnt int
		for i := 0; i+1 < n; i++ {
			// 如果 0和3能匹配，那么1， 2，也能和3匹配
			if nums[i+1]-nums[i] <= d {
				cnt++
				i++
			}
		}
		return cnt >= p
	}

	l, r := 0, nums[n-1]-nums[0]

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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
