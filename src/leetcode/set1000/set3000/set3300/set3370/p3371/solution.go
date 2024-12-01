package p3371

import "math/bits"

func smallestNumber(n int) int {
	h := bits.Len(uint(n))
	return 1<<h - 1
}
