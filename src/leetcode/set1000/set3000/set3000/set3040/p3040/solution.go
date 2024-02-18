package p3040

import "sort"

const oo = 1 << 30

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	mx := nums[n-1] + 1

	dp := make([]int, mx+1)

	for _, x := range nums {
		dp[x+1] = max(dp[x+1], dp[x]+1)
		dp[x] = max(dp[x], dp[x-1]+1)
	}

	var res int
	for _, v := range dp {
		res = max(res, v)
	}

	return res
}
