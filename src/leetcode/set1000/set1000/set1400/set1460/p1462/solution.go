package p1462

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := 0; i < n; i++ {
		for _, pre := range prerequisites {
			u, v := pre[0], pre[1]
			dp[i][v] = dp[i][v] || dp[i][u]
		}
	}

	res := make([]bool, len(queries))

	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		res[i] = dp[a][b]
	}

	return res
}
