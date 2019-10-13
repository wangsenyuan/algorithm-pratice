package p1223

const MOD = 1000000007

func dieSimulator(n int, rollMax []int) int {
	dp := make([][][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, 6)
		for j := 0; j < 6; j++ {
			dp[i][j] = make([]int64, 16)
		}
	}

	dp[0][0][0] = 1

	for i := 1; i <= n; i++ {
		for k := 0; k < 6; k++ {
			for j := 2; j <= rollMax[k]; j++ {
				dp[i][k][j] = dp[i-1][k][j-1]
			}

			var sum int64
			for j := 0; j < 6; j++ {
				if j == k && i > 1 {
					continue
				}
				for l := 0; l < 16; l++ {
					sum += dp[i-1][j][l]
					if sum >= MOD {
						sum -= MOD
					}
				}
			}

			dp[i][k][1] = sum
		}
	}
	var res int64
	for j := 0; j < 6; j++ {
		for l := 0; l < 16; l++ {
			res += dp[n][j][l]
			if res >= MOD {
				res -= MOD
			}
		}
	}
	return int(res)
}
