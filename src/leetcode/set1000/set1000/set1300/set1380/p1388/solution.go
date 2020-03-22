package p1388

func maxSizeSlices(slices []int) int {

	eat := func(arr []int, m int) int {
		n := len(arr)
		dp := make([][]int, n+2)
		for i := 0; i <= n+1; i++ {
			dp[i] = make([]int, m+1)
			for j := 1; j <= m; j++ {
				dp[i][j] = -10000000
			}
		}
		dp[0][0] = 0

		var res int

		for i := 2; i <= n+1; i++ {
			for j := 1; j <= m; j++ {
				dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+arr[i-2])
			}
			res = max(res, dp[i][m])
		}

		return res
	}

	n := len(slices)
	m := n / 3

	return max(eat(slices[:n-1], m), eat(slices[1:], m))
}

func maxSizeSlices1(slices []int) int {

	var dfs func(i, j, k, cycle int) int

	dfs = func(i, j, k, cycle int) int {
		if k == 1 {
			return maxOf(slices[i : j+1])
		}
		if j-i+1 < 2*k-1 {
			return -100000
		}

		a := slices[j] + dfs(i+cycle, j-2, k-1, 0)
		b := dfs(i, j-1, k, 0)

		return max(a, b)
	}

	return dfs(0, len(slices)-1, len(slices)/3, 1)
}

func maxOf(s []int) int {
	res := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] > res {
			res = s[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
