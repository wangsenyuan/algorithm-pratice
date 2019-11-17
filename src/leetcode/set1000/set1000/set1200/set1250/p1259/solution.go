package p1259

const MOD = 1000000007

func numberOfWays(n int) int {
	// 0... n - 1, 0 shake with 2, 3, 4,
	dp := make([]int, n+1)
	dp[0] = 1
	dp[2] = 1

	for x := 4; x <= n; x += 2 {
		// split x into a + b + 2
		for a := 0; a < x; a += 2 {
			b := x - a - 2
			tmp := mul(dp[a], dp[b])
			dp[x] += tmp
			if dp[x] >= MOD {
				dp[x] -= MOD
			}
		}
	}

	return dp[n]
}

func mul(a, b int) int {
	A := int64(a)
	B := int64(b)
	C := (A * B) % MOD
	return int(C)
}
