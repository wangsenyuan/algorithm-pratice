package p3091

import "sort"

func minOperations(k int) int {

	check := func(n int) bool {
		// 经过x次，所有的数字变成了x + 1, (x + 1) * (n - x + 1) >= k
		// let n + 1 = m
		// (x + 1) * (m - x) >= k
		//
		return n*n+4*(n+1-k) >= 0
	}

	return sort.Search(k, check)
}
