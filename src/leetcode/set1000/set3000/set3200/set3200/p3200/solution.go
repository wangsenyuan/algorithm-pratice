package p3200

import "slices"

func maximumLength(nums []int, k int) int {

	dp := make([]int, k)
	// f(sum) = when a[0] + a[1] % k = sum best answer
	f := func(sum int) int {
		for i := 0; i < k; i++ {
			dp[i] = 0
		}

		for _, num := range nums {
			r := num % k
			if r < sum {
				dp[r] = max(dp[r], dp[sum-r]+1)
			} else {
				dp[r] = max(dp[r], dp[(sum+k-r)%k]+1)
			}
		}
		return slices.Max(dp)
	}

	var res int

	for i := 0; i < k; i++ {
		res = max(res, f(i))
	}

	return res
}
