package p1898

import "sort"

func maximumRemovals(s string, p string, removable []int) int {
	back := make([]int, len(removable))
	check := func(k int) bool {
		copy(back, removable)
		sort.Ints(back[:k])
		var i, j int
		var x int
		for i < len(s) && j < len(p) {
			if x < k && back[x] == i {
				i++
				x++
				continue
			}
			// s[i] is not removed
			if s[i] == p[j] {
				j++
			}
			i++
		}
		return j == len(p)
	}

	if check(len(removable)) {
		return len(removable)
	}

	left, right := 1, len(removable)

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
