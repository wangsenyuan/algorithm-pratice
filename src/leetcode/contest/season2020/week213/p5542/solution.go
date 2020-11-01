package p5542

const MOD = 1000000007

func numWays(words []string, target string) int {
	m := len(words)
	n := len(words[0])

	cnt := make([][]int, n)

	for j := 0; j < n; j++ {
		cnt[j] = make([]int, 26)
		for i := 0; i < m; i++ {
			x := int(words[i][j] - 'a')
			cnt[j][x]++
		}
	}
	mm := len(target)

	dp := make([]int64, mm+1)
	dp[0] = 1

	for j := 0; j < n; j++ {
		for i := mm; i > 0; i-- {
			dp[i] += dp[i-1] * int64(cnt[j][int(target[i-1]-'a')]) % MOD
			dp[i] %= MOD
		}
	}

	return int(dp[mm])
}
