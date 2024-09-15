package p3290

const inf = 1 << 50

func maxScore(a []int, b []int) int64 {
	// len(a) = 4
	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = -inf
	}
	dp[0] = 0

	for _, num := range b {
		for j := 4; j > 0; j-- {
			dp[j] = max(dp[j], num*a[j-1]+dp[j-1])
		}
	}

	return int64(dp[4])
}
