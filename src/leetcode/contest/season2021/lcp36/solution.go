package lcp36

import (
	"sort"
)

func maxGroupNumber(tiles []int) int {
	sort.Ints(tiles)
	pp := make([]Pair, 0, 10)
	for i, j := 1, 0; i <= len(tiles); i++ {
		if i == len(tiles) || tiles[i] > tiles[i-1] {
			pp = append(pp, Pair{tiles[i-1], i - j})
			j = i
		}
	}
	// try to pair up
	dp := make([]int, 25)
	fp := make([]int, 25)

	for i := 1; i < 25; i++ {
		dp[i] = -1
	}

	prev := 0
	for i := 0; i < len(pp); i++ {
		cur := pp[i]
		if prev+1 != cur.first {
			for j := 1; j < 25; j++ {
				dp[j] = -1
			}
		}
		for j := 0; j < 25; j++ {
			fp[j] = -1
		}
		for j := 0; j < 25; j++ {
			if dp[j] < 0 {
				continue
			}
			a, b := j/5, j%5
			k := min(a, min(b, cur.second))
			for x := 0; x <= k; x++ {
				bb := b - x
				for c := 0; c <= min(4, cur.second-x); c++ {
					fp[bb*5+c] = max(fp[bb*5+c], dp[j]+x+(cur.second-c-x)/3)
				}
			}
		}
		copy(dp, fp)
		prev = cur.first
	}
	var best int
	for j := 0; j < 25; j++ {
		if dp[j] >= 0 {
			a, b := j/5, j%5
			best = max(best, dp[j]+a/3+b/3)
		}
	}
	return best
}

type Pair struct {
	first, second int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
