package p818

func racecar(target int) int {
	dp := make([]int, target+3)

	for i := 0; i < len(dp); i++ {
		dp[i] = 1 << 30
	}

	dp[0] = 0
	dp[1] = 1
	dp[2] = 4

	for t := 3; t <= target; t++ {
		k := 32 - numOfLeadingZeros(t)
		if t == 1<<uint(k)-1 {
			dp[t] = k
			continue
		}

		for j := 0; j < k-1; j++ {
			dp[t] = min(dp[t], dp[t-(1<<uint(k-1))+1<<uint(j)]+k-1+j+2)
		}
		if 1<<uint(k)-1-t < t {
			dp[t] = min(dp[t], dp[1<<uint(k)-1-t]+k+1)
		}
	}

	return dp[target]
}

func numOfLeadingZeros(x int) int {
	var i int

	for x > 0 {
		x >>= 1
		i++
	}

	return 32 - i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
