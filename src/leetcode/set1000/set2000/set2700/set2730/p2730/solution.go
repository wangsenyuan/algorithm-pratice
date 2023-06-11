package p2730

import "sort"

const MOD = 1_000_000_007

// 8920312451
func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modVal(x int64) int {
	x %= MOD
	if x < 0 {
		x += MOD
	}
	return int(x)
}

func sub(a, b int) int {
	return add(a, MOD-b)
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func sumDistance(nums []int, s string, d int) int {
	n := len(nums)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			arr[i] = nums[i] - d
		} else {
			arr[i] = nums[i] + d
		}
	}
	sort.Ints(arr)

	dist := make([]int64, n-1)
	for i := 0; i+1 < n; i++ {
		dist[i] = int64(arr[i+1]) - int64(arr[i])
	}

	calc := func(x, y, z int) int {
		return sub(mul(x, y), z)
	}

	// sum[i] - (sum[i-1] - sum[i-2] - ... sum[0])
	var sum_of_sum int
	var sum int
	var res int
	for i := 0; i < n-1; i++ {
		sum = add(sum, modVal(dist[i]))
		res = add(res, calc(sum, i+1, sum_of_sum))
		sum_of_sum = add(sum_of_sum, sum)
	}

	return res
}
