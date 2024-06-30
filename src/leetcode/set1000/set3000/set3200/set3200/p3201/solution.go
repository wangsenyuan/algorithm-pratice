package p3201

import "slices"

func maximumLength(nums []int) int {
	dp := make([]int, 2)
	// f(sum) = when a[0] + a[1] % k = sum best answer
	f := func(sum int) int {
		for i := 0; i < 2; i++ {
			dp[i] = 0
		}

		for _, num := range nums {
			r := num % 2
			if r < sum {
				dp[r] = max(dp[r], dp[sum-r]+1)
			} else {
				dp[r] = max(dp[r], dp[(sum+2-r)%2]+1)
			}
		}
		return slices.Max(dp)
	}

	return max(f(0), f(1))
}
