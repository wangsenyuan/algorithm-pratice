package p1220

const MOD = 1000000007

func countVowelPermutation(n int) int {
	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}
	fp := make([]int, 5)
	// 'a', 'e', 'i', 'o', 'u
	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			fp[j] = 0
		}
		fp[1] = dp[0]
		fp[0] = dp[1]
		fp[2] = dp[1]
		fp[0] += dp[2]
		fp[1] += dp[2]
		fp[3] += dp[2]
		fp[4] += dp[2]
		fp[2] += dp[3]
		fp[4] += dp[3]
		fp[0] += dp[4]

		for j := 0; j < 5; j++ {
			dp[j] = fp[j] % MOD
		}
	}

	var sum int

	for j := 0; j < 5; j++ {
		sum += dp[j]
	}

	return sum % MOD
}
