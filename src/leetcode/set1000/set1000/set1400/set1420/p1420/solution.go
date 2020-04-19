package p1420

const MOD = 1000000007

func pow(a, b int) int64 {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R *= A
			R %= MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}

func numOfArrays(n int, M int, K int) int {
	dp := make([][][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, M+1)
		for j := 0; j <= M; j++ {
			dp[i][j] = make([]int64, K+1)
		}
	}

	for m := 1; m <= M; m++ {
		dp[1][m][1] = 1
	}

	for i := 2; i <= n; i++ {
		// try to compute dp[i][m][k]
		for m := 1; m <= M; m++ {
			// try to put m at pos ii
			for k := 1; k <= K; k++ {
				dp[i][m][k] = (int64(m) * dp[i-1][m][k]) % MOD

				for x := 1; x < m; x++ {
					dp[i][m][k] += dp[i-1][x][k-1]
					if dp[i][m][k] >= MOD {
						dp[i][m][k] -= MOD
					}
				}
			}
		}
	}

	var res int64

	for m := 1; m <= M; m++ {
		res += dp[n][m][K]
		if res >= MOD {
			res -= MOD
		}
	}

	return int(res)
}
