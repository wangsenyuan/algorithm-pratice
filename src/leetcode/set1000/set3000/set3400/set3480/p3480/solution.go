package p3480

import "slices"

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	group := make([][]int, n+1)
	for _, cur := range conflictingPairs {
		a, b := cur[0], cur[1]
		if a > b {
			a, b = b, a
		}
		group[b] = append(group[b], a)
	}

	first, second := 0, 0
	extra := make([]int, n+1)
	var ans int

	for b := 1; b <= n; b++ {
		for _, a := range group[b] {
			if a >= first {
				second = first
				first = a
			} else if a >= second {
				second = a
			}
		}
		ans += b - first
		extra[first] += first - second
	}
	return int64(ans + slices.Max(extra))
}
