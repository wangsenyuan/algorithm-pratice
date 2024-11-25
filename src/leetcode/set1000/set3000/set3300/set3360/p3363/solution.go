package p3363

const inf = 1 << 30

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)

	// 三个区域是独立的
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -inf
		}
	}

	get := func(r int, c int) int {
		if r < 0 || r >= n || c < 0 || c >= n {
			return -inf
		}
		return dp[r][c]
	}
	dp[0][n-1] = fruits[0][n-1]
	for i := 1; i+1 < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(get(i-1, j-1), get(i-1, j), get(i-1, j+1)) + fruits[i][j]
		}
	}
	res := dp[n-2][n-1]

	dp[n-1][0] = fruits[n-1][0]

	for j := 1; j < n; j++ {
		for i := j + 1; i < n; i++ {
			dp[i][j] = max(get(i-1, j-1), get(i, j-1), get(i+1, j-1)) + fruits[i][j]
		}
	}

	res += dp[n-1][n-2]

	for i := 0; i < n; i++ {
		res += fruits[i][i]
	}

	return res
}
