package p2942

func findMaximumLength(nums []int) int {
	n := len(nums)

	dp := make([]int, n+1)
	sum := make([]int, n+1)
	last := make([]int, n+1)
	que := make([]int, n+1)
	var front, tail int
	que[front] = 0
	front++
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]

		for tail+1 < front && sum[que[tail+1]]+last[que[tail+1]] <= sum[i] {
			tail++
		}

		dp[i] = dp[que[tail]] + 1
		last[i] = sum[i] - sum[que[tail]]

		for front > tail && sum[que[front-1]]+last[que[front-1]] >= sum[i]+last[i] {
			front--
		}
		que[front] = i
		front++
	}

	return dp[n]
}
