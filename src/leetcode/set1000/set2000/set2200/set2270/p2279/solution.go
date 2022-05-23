package p2279

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	n := len(capacity)
	left := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = capacity[i] - rocks[i]
	}
	sort.Ints(left)
	var sum int
	for i := 0; i < n; i++ {
		sum += left[i]
		if sum > additionalRocks {
			return i
		}
	}
	return n
}
