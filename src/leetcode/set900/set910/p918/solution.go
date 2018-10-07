package p918

import (
	"math"
)

func maxSubarraySumCircular(A []int) int {
	n := len(A)

	res := math.MinInt32
	var sum int
	for i := 0; i < n; i++ {
		sum += A[i]
		if res < sum {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	sum = 0

	f := make([]int, n)
	for i := 0; i < n; i++ {
		sum += A[i]
		f[i] = sum
		if i > 0 && f[i-1] > sum {
			f[i] = f[i-1]
		}
	}

	sum = 0
	fromRight := 0

	for i := n - 1; i > 0; i-- {
		sum += A[i]
		if sum > fromRight {
			fromRight = sum
		}
		if f[i-1]+fromRight > res {
			res = f[i-1] + fromRight
		}
	}
	return res
}
