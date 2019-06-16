package p910

import (
	"sort"
)

func smallestRangeII(A []int, K int) int {
	sort.Ints(A)
	n := len(A)

	ans := A[n-1] - A[0]

	a := A[0] + K
	b := A[n-1] - K
	for i := 0; i < n-1; i++ {
		// before i, all increase by K, after i, all decrease by K
		x := A[i] + K
		y := A[i+1] - K

		c := max(b, x)
		d := min(a, y)
		if c-d < ans {
			ans = c - d
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
