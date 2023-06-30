package lcp77

import "sort"

func runeReserve(runes []int) int {
	sort.Ints(runes)

	var best int
	n := len(runes)

	for i := 0; i < n; {
		j := i
		i++
		for i < n && runes[i]-runes[i-1] <= 1 {
			i++
		}
		best = max(best, i-j)
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
