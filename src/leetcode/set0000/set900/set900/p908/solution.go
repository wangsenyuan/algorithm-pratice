package p908

import "math"

func smallestRangeI(A []int, K int) int {
	x := math.MaxInt32
	y := math.MinInt32

	for i := 0; i < len(A); i++ {
		if A[i] < x {
			x = A[i]
		}
		if A[i] > y {
			y = A[i]
		}
	}

	// x, y is the min, max value
	x += K
	y -= K
	if y <= x {
		return 0
	}
	return y - x
}
