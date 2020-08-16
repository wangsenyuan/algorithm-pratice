package p1552

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	check := func(x int) bool {
		cnt := m
		prev := position[0]
		cnt--
		for i := 1; i < n && cnt > 0; i++ {
			if position[i]-prev >= x {
				// have to put one ball before current i
				cnt--
				prev = position[i]
			}
		}
		return cnt == 0
	}

	var left int
	var right = position[n-1] - position[0] + 1

	for left < right {
		mid := (left + right) / 2
		if !check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}
