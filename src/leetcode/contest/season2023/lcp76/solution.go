package lcp76

func getSchemeCount(n int, m int, board []string) int64 {
	// 两颗异色棋子在同一行或者同一列
	// 两颗异色棋子之间恰好只有一颗棋子
	if n < m {
		n, m = m, n
		board = transpose(board)
	}

	// 0 for no color
	// 1 for red only
	// 2 for black only
	// 3 for RB
	// 4 for BR
	// 5 for RR
	// 6 for BB

	T := 7

	findNextState := func(cur int, x byte) int {
		if x == '.' {
			// no change
			return cur
		}
		if cur == 0 {
			if x == 'R' {
				return 1
			}
			return 2
		}

		if cur == 4 || cur == 5 {
			cur = 1
		} else if cur == 3 || cur == 6 {
			cur = 2
		}

		if cur == 1 {
			if x == 'R' {
				return 5
			}
			return 3
		}
		if x == 'R' {
			return 4
		}
		return 6
	}

	check := func(r int, c int, x byte, row int, state int) bool {
		if x == '.' {
			return true
		}
		col := state
		for i := m - 1; i > c; i-- {
			col /= T
		}
		col %= T

		if x == 'R' {
			// BR BB not work
			if row == 4 || row == 6 || col == 4 || col == 6 {
				return false
			}
		} else {
			// RB or RR not work
			if row == 3 || row == 5 || col == 3 || col == 5 {
				return false
			}
		}

		return true
	}

	buf := make([]byte, m)

	getNextState := func(state int) int {
		var next int
		var t = 1
		for i := m - 1; i >= 0; i-- {
			cur := state % T
			x := buf[i]
			tmp := findNextState(cur, x)
			next = tmp*t + next
			t *= T
			state /= T
		}

		return next
	}

	dp := make(map[int]int64)
	fp := make(map[int]int64)
	dp[0] = 1

	var dfs func(r int, c int, row int, state int)

	dfs = func(r int, c int, row int, state int) {
		if c == m {
			next := getNextState(state)
			fp[next] += dp[state]
			return
		}
		buf[c] = board[r][c]

		if board[r][c] != '?' {
			if check(r, c, board[r][c], row, state) {
				dfs(r, c+1, findNextState(row, board[r][c]), state)
			}
			return
		}
		buf[c] = '.'
		dfs(r, c+1, findNextState(row, '.'), state)

		if check(r, c, 'R', row, state) {
			buf[c] = 'R'
			dfs(r, c+1, findNextState(row, 'R'), state)
		}
		if check(r, c, 'B', row, state) {
			buf[c] = 'B'
			dfs(r, c+1, findNextState(row, 'B'), state)
		}
	}

	for r := 0; r < n; r++ {
		for state := range dp {
			dfs(r, 0, 0, state)
		}
		dp = fp
		fp = make(map[int]int64)
	}

	var res int64

	for _, v := range dp {
		res += v
	}
	return res
}

func transpose(board []string) []string {
	n := len(board)
	m := len(board[0])
	buf := make([][]byte, m)
	for i := 0; i < m; i++ {
		buf[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			buf[j][i] = board[i][j]
		}
	}
	res := make([]string, m)
	for i := 0; i < m; i++ {
		res[i] = string(buf[i])
	}
	return res
}
