package p1191

import "math"

const MOD = 1e9 + 7

func kConcatenationMaxSum(arr []int, k int) int {
	n := len(arr)
	var best int64
	if k == 1 {
		best = findMaxSumSum(n, func(i int) int {
			return arr[i]
		})
		return int(best % MOD)
	}

	// then sum have to be in the range arr + arr
	best = findMaxSumSum(2*n, func(i int) int {
		return arr[i%n]
	})

	best = max(0, best)

	if k == 2 {
		return int(best % MOD)
	}

	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(arr[i])
	}

	if sum <= 0 {
		return int(best % MOD)
	}

	best = max(best, sum*int64(k))

	var a int64
	sum = 0
	for i := 0; i < n; i++ {
		sum += int64(arr[i])
		a = max(a, sum)
	}
	var b int64
	sum = 0
	for i := n - 1; i >= 0; i-- {
		sum += int64(arr[i])
		b = max(b, sum)
	}

	best = max(best, a+sum*int64(k-1))
	best = max(best, b+sum*int64(k-1))
	best = max(best, a+b+sum*int64(k-2))

	return int(best % MOD)
}

func findMaxSumSum(n int, ele func(int) int) int64 {
	var sum int64
	var best int64 = math.MinInt64
	for i := 0; i < n; i++ {
		sum += int64(ele(i))
		best = max(best, sum)
		sum = max(0, sum)
	}
	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
