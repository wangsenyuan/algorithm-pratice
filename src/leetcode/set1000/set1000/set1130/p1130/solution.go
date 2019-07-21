package p1130

const INF = 1 << 30

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	fp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
		fp[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
		dp[i][0] = 0
		dp[i][1] = 0
		fp[i][1] = arr[i]
	}

	for k := 2; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			x := arr[i]
			ans := INF
			for u := 1; u < k; u++ {
				v := k - u
				tmp := fp[i][u]*fp[i+u][v] + dp[i][u] + dp[i+u][v]
				ans = min(ans, tmp)
				if x < arr[i+u] {
					x = arr[i+u]
				}
			}
			dp[i][k] = ans
			fp[i][k] = x
		}
	}

	return dp[0][n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
