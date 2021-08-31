package p1984

import "sort"

func minimumDifference(A []int, k int) int {
	sort.Ints(A)

	best := 1 << 30
	for i := k - 1; i < len(A); i++ {
		best = min(best, A[i]-A[i-k+1])
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
