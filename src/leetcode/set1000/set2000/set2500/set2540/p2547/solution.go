package p2547

func minCost(nums []int, k int) int {
	n := len(nums)
	// 要连续的子数组
	score := make([][]int, n)
	for i := 0; i < n; i++ {
		score[i] = make([]int, n)
		cnt := make(map[int]int)
		var ones int
		for j := i; j < n; j++ {
			cnt[nums[j]]++
			if cnt[nums[j]] == 1 {
				ones++
			} else if cnt[nums[j]] == 2 {
				ones--
			}
			score[i][j] = j - i + 1 - ones
		}
	}

	// dp[i]
	dp := make([]int, n)
	// 具体分了几个不重要
	//dp[0] = k
	for i := 0; i < n; i++ {
		dp[i] = score[0][i] + k
		for j := i - 1; j >= 0; j-- {
			dp[i] = min(dp[i], dp[j]+score[j+1][i]+k)
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
