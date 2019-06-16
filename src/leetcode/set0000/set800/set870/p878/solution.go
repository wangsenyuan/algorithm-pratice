package p878

import "math"

const MOD = 1000000007

func nthMagicalNumber(n int, a int, b int) int {
	A := int64(a)
	B := int64(b)
	C := A * B / gcd(A, B)
	N := int64(n)
	left, right := int64(0), int64(math.MaxInt64)
	for left < right {
		mid := (left + right) >> 1
		num := mid/A + mid/B - mid/C
		if num >= N {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int(left % MOD)
}

func gcd(A, B int64) int64 {
	for B > 0 {
		A, B = B, A%B
	}
	return A
}
