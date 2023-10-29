package lcp11

import "sort"

func expectNumber(scores []int) int {
	sort.Ints(scores)

	n := len(scores)
	var res int

	for i := 1; i <= n; i++ {
		if i == n || scores[i] > scores[i-1] {
			res++
		}
	}

	return res
}
