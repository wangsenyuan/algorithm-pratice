package lcr137

func articleMatch(s string, p string) bool {
	// .匹配任何一个字符
	// *匹配0个或者1个字符
	n := len(s)
	m := len(p)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	ok := func(i int, j int) bool {
		if i >= n || j >= m {
			return false
		}
		return s[i] == p[j] || p[j] == '.'
	}

	var dfs func(i int, j int) bool

	dfs = func(i int, j int) bool {
		if j == m {
			return i == n
		}

		if dp[i][j] != 0 {
			return dp[i][j] > 0
		}
		var res bool
		if j+1 < m && p[j+1] == '*' {
			// 前面的字符不匹配
			res = dfs(i, j+2)
			// 或者匹配多次
			if !res && ok(i, j) {
				res = dfs(i+1, j)
			}
		}
		if !res && ok(i, j) {
			res = dfs(i+1, j+1)
		}
		if res {
			dp[i][j] = 1
		} else {
			dp[i][j] = -1
		}
		return res
	}

	return dfs(0, 0)
}
