package p1751

import "sort"

func maxValue(events [][]int, k int) int {
	n := len(events)

	sort.Slice(events, func(i, j int) bool {
		a, b := events[i], events[j]
		return a[1] < b[1]
	})

	// dp[i][j]  = 前i个event，选择最多j个时的最大值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			dp[i][j] = -1
		}
	}
	var best int
	for i := 0; i < n; i++ {
		cur := events[i]
		j := sort.Search(i, func(j int) bool {
			return events[j][1] >= cur[0]
		})
		//events[j-1][1] < cur[0]
		if j == 0 {
			dp[i][1] = cur[2]
		} else {
			for x := 0; x < k; x++ {
				dp[i][x+1] = dp[j-1][x] + cur[2]
			}
		}
		for x := 1; x <= k; x++ {
			dp[i][x] = max(dp[i][x], dp[i][x-1])
		}
		if i > 0 {
			for x := 1; x <= k; x++ {
				dp[i][x] = max(dp[i][x], dp[i-1][x])
			}
		}

		best = max(best, dp[i][k])
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
