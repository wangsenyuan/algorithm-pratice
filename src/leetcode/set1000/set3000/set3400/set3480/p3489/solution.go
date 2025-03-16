package p3489

import "slices"

func minZeroArray(nums []int, queries [][]int) int {
	m := len(queries)
	mx := slices.Max(nums)
	dp := make([]bool, mx+1)
	get := func(pos int, u int) int {
		if u == 0 {
			return 0
		}
		clear(dp)
		dp[0] = true
		for i, cur := range queries {
			l, r, v := cur[0], cur[1], cur[2]
			if r < pos || pos < l {
				continue
			}
			for j := mx; j >= v; j-- {
				if !dp[j] {
					dp[j] = dp[j-v]
				}
			}
			if dp[u] {
				return i + 1
			}
		}

		return m + 1
	}

	var best int
	for i, x := range nums {
		tmp := get(i, x)
		if tmp == m+1 {
			return -1
		}
		best = max(best, tmp)
	}

	return best
}
