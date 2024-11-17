package p3356

import "sort"

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	diff := make([]int, n+1)
	check := func(k int) bool {
		clear(diff)
		for i := 0; i < k; i++ {
			l, r, v := queries[i][0], queries[i][1], queries[i][2]
			diff[l] += v
			diff[r+1] -= v
		}

		for i := 1; i < n; i++ {
			diff[i] += diff[i-1]
		}

		for i := 0; i < n; i++ {
			if diff[i] < nums[i] {
				return false
			}
		}
		return true
	}

	if !check(len(queries)) {
		return -1
	}

	return sort.Search(len(queries), check)
}
