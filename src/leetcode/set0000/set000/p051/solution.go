package p051

func solveNQueens(n int) [][]string {
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			buf[i][j] = '.'
		}
	}

	create := func() []string {
		res := make([]string, n)
		for i := 0; i < n; i++ {
			res[i] = string(buf[i])
		}
		return res
	}

	check := func(r int, c int) bool {
		// 通一行不需要检查
		for i := r - 1; i >= 0; i-- {
			if buf[i][c] == 'Q' {
				return false
			}
			j := c - (r - i)
			if j >= 0 && buf[i][j] == 'Q' {
				return false
			}
			j = c + (r - i)
			if j < n && buf[i][j] == 'Q' {
				return false
			}
		}
		return true
	}
	var dfs func(r int)

	var ans [][]string

	dfs = func(r int) {
		if r == n {
			ans = append(ans, create())
			return
		}
		for c := 0; c < n; c++ {
			if check(r, c) {
				buf[r][c] = 'Q'
				dfs(r + 1)
				buf[r][c] = '.'
			}
		}
	}

	dfs(0)

	return ans
}
