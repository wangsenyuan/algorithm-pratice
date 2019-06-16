package p940

const MOD = 1e9 + 7

func distinctSubseqII(S string) int {
	n := len(S)

	dp := make([]int64, n+1)
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		last[i] = -1
	}
	dp[0] = 1
	for i := 0; i < n; i++ {
		dp[i+1] = 2 * dp[i] % MOD
		if last[S[i]-'a'] >= 0 {
			dp[i+1] -= dp[last[S[i]-'a']]
			if dp[i+1] < 0 {
				dp[i+1] += MOD
			}
		}
		last[S[i]-'a'] = i
	}

	dp[n]--
	if dp[n] < 0 {
		dp[n] += MOD
	}

	return int(dp[n])
}
