package p730

const MOD = 1000000007

func countPalindromicSubsequences(S string) int {
	n := len(S)

	X := make([]int, n)
	for i := 0; i < n; i++ {
		X[i] = int(S[i] - 'a')
	}

	dp := make([][][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 4)
		}
	}

	for l := 1; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			for x := 0; x < 4; x++ {
				var ans int
				if l == 1 {
					if X[i] == x {
						ans = 1
					}
				} else {
					if X[i] != x {
						ans = dp[1][i+1][x]
					} else if X[j] != x {
						ans = dp[1][i][x]
					} else if l == 2 {
						ans = 2
					} else {
						ans = 2
						for y := 0; y < 4; y++ {
							ans += dp[0][i+1][y]
							ans %= MOD
						}
					}
				}
				dp[2][i][x] = ans
			}
		}
		for i := 0; i < 2; i++ {
			for j := 0; j < n; j++ {
				for x := 0; x < 4; x++ {
					dp[i][j][x] = dp[i+1][j][x]
				}
			}
		}
	}

	var ret int

	for x := 0; x < 4; x++ {
		ret += dp[2][0][x]
		ret %= MOD
	}

	return ret
}
