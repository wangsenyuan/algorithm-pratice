package p2710

const INF = 1 << 60

func minimumCost(s string) int64 {
	dp := process(s)
	fp := process(reverse(s))
	n := len(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		fp[i], fp[j] = fp[j], fp[i]
	}

	best := min(dp[n-1][0], dp[n-1][1])

	for i := 0; i+1 < n; i++ {
		for j := 0; j < 2; j++ {
			tmp := dp[i][j] + fp[i+1][j]
			best = min(best, tmp)
		}
	}
	return best
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func process(s string) [][]int64 {
	n := len(s)

	// dp[i][0] = 将pref i都变成0的cost
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][int(s[0]-'0')] = 0
	dp[0][1-int(s[0]-'0')] = 1

	for i := 1; i < n; i++ {
		x := int(s[i] - '0')
		// dp[i][0] = dp[i-1][0], i + dp[i-1][1]
		// 表示，如果前面都是0了，不需要翻转，或者将前面的1都翻转后，再增加0
		dp[i][x] = min(dp[i-1][x], int64(i)+dp[i-1][1^x])
		// 表示，如果前面都是0，先加上去，然后翻转i+1个，或者前面时1，先翻转成0，再全部翻转成1
		// dp[i][1] = dp[i-1][0] + i + 1, dp[i-1][1] + i + i + 1
		dp[i][1^x] = min(dp[i-1][x]+int64(i+1), dp[i-1][1^x]+int64(2*i+1))
	}

	return dp
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
