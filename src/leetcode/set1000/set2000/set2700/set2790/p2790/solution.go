package p2790

import "sort"

func maxIncreasingGroups(usageLimits []int) int {
	sort.Ints(usageLimits)

	reverse(usageLimits)

	check := func(n int) bool {
		var gap int
		for _, x := range usageLimits {
			gap = min(gap+x-n, 0)
			if n > 0 {
				n--
			}
		}
		return gap >= 0
	}

	l, r := 0, len(usageLimits)+1
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
