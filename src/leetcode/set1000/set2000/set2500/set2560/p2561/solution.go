package p2561

const INF = 1 << 30

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	// dp[i][1] 等于满足条件，一个segment中获得的结果
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = -INF
	}
	var best int
	for i := 0; i < n; i++ {
		// 如果以prize[i]的位置为结尾
		cur := prizePositions[i]
		j := search(i+1, func(j int) bool {
			return prizePositions[j] >= cur-k
		})
		best = max(best, dp[j]+i-j+1)
		dp[i+1] = max(dp[i], i-j+1)
	}

	return best
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
