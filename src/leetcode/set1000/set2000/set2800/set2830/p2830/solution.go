package p2830

func maximizeTheProfit(n int, offers [][]int) int {

	at := make([][]int, n)

	for i, cur := range offers {
		end := cur[1]
		if len(at[end]) == 0 {
			at[end] = make([]int, 0, 1)
		}
		at[end] = append(at[end], i)
	}
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		for _, j := range at[i] {
			start, cost := offers[j][0], offers[j][2]
			if start > 0 {
				dp[i] = max(dp[i], dp[start-1]+cost)
			} else {
				dp[i] = max(dp[i], cost)
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
