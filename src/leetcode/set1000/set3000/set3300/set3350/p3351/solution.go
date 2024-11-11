package p3351

import "slices"

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func sumOfGoodSubsequences(nums []int) int {
	x := slices.Max(nums)

	dp := make([]int, x+2)

	n := len(nums)

	fp := make([]int, n)

	for i, num := range nums {
		tmp := dp[num+1]
		if num > 0 {
			tmp = add(tmp, dp[num-1])
		}
		fp[i] = add(tmp, 1)
		dp[num] = add(dp[num], fp[i])
	}

	clear(dp)

	var res int

	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		tmp := dp[num+1]
		if num > 0 {
			tmp = add(tmp, dp[num-1])
		}
		tmp = add(tmp, 1)
		dp[num] = add(dp[num], tmp)

		tmp = mul(tmp, fp[i])
		res = add(res, mul(tmp, num))
	}

	return res
}
