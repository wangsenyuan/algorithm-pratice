package p983

const INF = 1 << 20

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	if n == 0 {
		return 0
	}
	//dp[i] = min cost before days i
	dp := make([]int, 366)
	for i := 0; i < 366; i++ {
		dp[i] = INF
	}

	var j int
	dp[0] = 0
	for i := 1; i < 366; i++ {
		if j == len(days) || i < days[j] {
			dp[i] = min(dp[i-1], dp[i])
			continue
		}
		//i == dp[j]
		// buy ticket one
		dp[i] = min(dp[i], dp[i-1]+costs[0])
		// buy ticket two
		for k := 0; k < 7 && i+k < 366; k++ {
			dp[i+k] = min(dp[i+k], dp[i-1]+costs[1])
		}

		for k := 0; k < 30 && i+k < 366; k++ {
			dp[i+k] = min(dp[i+k], dp[i-1]+costs[2])
		}
		j++
	}

	return dp[365]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
