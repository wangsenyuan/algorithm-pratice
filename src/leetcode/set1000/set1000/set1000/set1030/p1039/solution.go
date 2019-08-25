package p1039

const INF = 1 << 30

func minScoreTriangulation(A []int) int {
	n := len(A)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for k := 2; k < n; k++ {
		for i := 0; i+k < n; i++ {
			j := i + k
			dp[i][j] = INF
			for u := i + 1; u < j; u++ {
				dp[i][j] = min(dp[i][j], dp[i][u]+dp[u][j]+A[i]*A[j]*A[u])
			}
		}
	}
	return dp[0][n-1]
}

func minScoreTriangulation1(A []int) int {
	n := len(A)

	// dp[i][j] = starting from point i, with j consective points
	dp := make([][][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = INF
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j][2] = A[i] * A[j] * A[(j+1)%n]
		}
	}

	for k := 3; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				u := (j + k - 1) % n
				if u < i || u > j {
					tmp := A[i]*A[j]*A[u] + dp[j][(j+1)%n][k-1]

					for x := 2; x < k; x++ {
						v := (j + x - 1) % n
						tmp = min(tmp, dp[i][j][x]+dp[i][v][k-x+1])
					}
					dp[i][j][k] = tmp
				}
			}
		}
	}
	return dp[0][1][n-1]
}

func inBetween(a, b, c int) bool {
	if c >= a && c <= b {
		return true
	}
	return false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
