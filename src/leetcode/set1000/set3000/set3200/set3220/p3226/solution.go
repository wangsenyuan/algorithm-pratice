package p3226

import "math/bits"

func minChanges(n int, k int) int {
	if n&k != k {
		return -1
	}
	n ^= k
	return bits.OnesCount32(uint32(n))
}
