package p1278

const INF = 1 << 20

func palindromePartition(s string, k int) int {
	n := len(s)

	fp := make([][]int, n)

	for i := 0; i < n; i++ {
		fp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[j] == s[i] {
				fp[j][i] = fp[j+1][i-1]
			} else {
				fp[j][i] = fp[j+1][i-1] + 1
			}
		}
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {

		for j := i; j >= 0; j-- {
			r := fp[j][i]

			for x := 0; x < k; x++ {
				if dp[j][x] < INF {
					dp[i+1][x+1] = min(dp[i+1][x+1], dp[j][x]+r)
				}
			}
		}
	}

	return dp[n][k]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
