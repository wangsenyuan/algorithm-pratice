package p416

import (
	"sort"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}
	half := sum / 2
	dp := make([]bool, half+1)
	dp[0] = true

	for _, cur := range nums {
		for j := half; j >= cur; j-- {
			dp[j] = dp[j] || dp[j-cur]
		}
	}

	return dp[half]
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}
	n := len(nums)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, 2*sum+1)
	}
	dp[0][sum] = true

	for i := 1; i <= n; i++ {
		cur := nums[i-1]
		for diff := 0; diff <= 2*sum; diff++ {
			if dp[i-1][diff] {
				dp[i][diff+cur] = true
				if diff-cur >= 0 {
					dp[i][diff-cur] = true
				}
			}
		}
	}

	return dp[n][sum]
}

func canPartition1(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	half := sum / 2
	sort.Ints(nums)

	var pick func(i int, sum int) bool

	pick = func(i int, sum int) bool {
		if sum == half {
			return true
		}
		if i < 0 {
			return false
		}

		if i*nums[i]+sum < half {
			return false
		}

		if sum+nums[i] <= half && pick(i-1, sum+nums[i]) {
			return true
		}
		return pick(i-1, sum)
	}

	return pick(len(nums)-1, 0)
}
