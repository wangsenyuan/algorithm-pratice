package p1846

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	var best int
	prev := 0
	for i := 0; i < n; i++ {
		prev = min(prev+1, arr[i])
		best = max(best, prev)
	}

	return best
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
