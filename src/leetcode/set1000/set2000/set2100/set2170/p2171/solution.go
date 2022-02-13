package p2171

import (
	"math"
	"sort"
)

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var sum int64
	for _, num := range beans {
		sum += int64(num)
	}
	var best int64 = math.MaxInt64
	var cur int64
	n := len(beans)
	for i := 0; i < n; i++ {
		// if take beans[i] as the base
		tmp := int64(beans[i]) * int64(n-i)
		tmp = sum - tmp
		best = min(best, cur+tmp)
		sum -= int64(beans[i])
		cur += int64(beans[i])
	}

	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
