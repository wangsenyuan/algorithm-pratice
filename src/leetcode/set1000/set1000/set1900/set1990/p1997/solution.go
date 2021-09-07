package p1997

const M = 1000000007

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	// cnt[i] % 2 == 0 时，下一个房间
	// else, nextVisit[i]
	// dp[i] 第一次进入i的天数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = 1
	// nextVisit[0] == 0
	for i := 1; i < n; i++ {
		dp[i][0] = modAdd(dp[i-1][1], 1)
		j := nextVisit[i]
		dp[i][1] = modAdd(dp[i][0], 1)

		if j < i {
			dp[i][1] = modAdd(dp[i][1], modAdd(modSub(dp[i-1][1], dp[j][0]), 1))
		}
	}

	return dp[n-1][0]
}

func modAdd(a, b int) int {
	a += b
	if a >= M {
		a -= M
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, M-b)
}
