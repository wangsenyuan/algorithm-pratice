package p2274

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {

	if len(special) == 0 {
		return top - bottom + 1
	}

	sort.Ints(special)

	var best int

	n := len(special)

	for i := 0; i < n; i++ {
		best = max(best, special[i]-bottom)
		bottom = special[i] + 1
	}

	best = max(best, top+1-bottom)

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
