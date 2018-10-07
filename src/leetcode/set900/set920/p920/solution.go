package p920

const MOD = 1e9 + 7

func numMusicPlaylists(N int, L int, K int) int {
	dp := make([][]int64, L+1)
	for i := 0; i <= L; i++ {
		dp[i] = make([]int64, N+1)
	}

	dp[0][0] = 1
	for i := 1; i <= L; i++ {
		for j := 1; j <= N; j++ {
			dp[i][j] += dp[i-1][j-1] * int64(N-j+1)
			dp[i][j] += dp[i-1][j] * int64(max(j-K, 0))
			dp[i][j] %= MOD
		}
	}

	return int(dp[L][N])
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
