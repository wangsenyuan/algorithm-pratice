package p1492

import "sort"

func kthFactor(n int, k int) int {
	N := int64(n)

	factors := make([]int, 0, 10)
	var th int
	for x := int64(1); x*x <= N; x++ {
		if N%x == 0 {
			th++
			if th == k {
				return int(x)
			}
			factors = append(factors, int(x))
			y := N / x
			if y != x {
				factors = append(factors, int(y))
			}
		}
	}

	sort.Ints(factors)

	if len(factors) < k {
		return -1
	}
	return factors[k-1]
}
