package p875

import "math"

func minEatingSpeed(piles []int, H int) int {

	check := func(K int) bool {
		var took int64
		for i := 0; i < len(piles); i++ {
			took += (int64(piles[i]) + int64(K) - 1) / int64(K)
		}
		return took <= int64(H)
	}

	left, right := 1, math.MaxInt32-20

	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
