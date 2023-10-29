package p2914

const inf = 1 << 30

func lengthOfLongestSubsequence(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = -inf
	}
	dp[0] = 0

	for _, num := range nums {
		for x := target; x >= num; x-- {
			dp[x] = max(dp[x], dp[x-num]+1)
		}
	}
	res := dp[target]
	return max(-1, res)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
