package p1246

const INF = 1 << 20

func minimumMoves(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j <= i; j++ {
			// just do one removal
			dp[i][j] = 1
		}
		for j := i + 1; j < n; j++ {
			dp[i][j] = INF
		}
	}

	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			// dp[i][j] = dp[i+1][j-1] if arr[i] == arr[j]
			if arr[i] == arr[j] {
				// an extension remove from [i+1...j-1]
				dp[i][j] = dp[i+1][j-1]
			} else {
				// dp[i][j] = dp[i+1][j] + 1
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}

	return dp[0][n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
