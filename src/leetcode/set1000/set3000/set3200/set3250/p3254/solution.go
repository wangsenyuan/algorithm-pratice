package p3254

const inf = 1 << 60

func maximumValueSum(board [][]int) int64 {
	n := len(board)
	m := len(board[0])

	type pair struct {
		first  int
		second int
	}

	dp := make([][][][]pair, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]pair, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([][]pair, 4)
			for k := 0; k < 4; k++ {
				dp[i][j][k] = make([]pair, 2)
				for u := 0; u < 2; u++ {
					dp[i][j][k][u] = pair{-inf, -1}
				}
			}
		}
	}
	col := make([][]pair, m)
	for i := 0; i < m; i++ {
		col[i] = make([]pair, 2)
		for j := 0; j < 2; j++ {
			col[i][j] = pair{-inf, -1}
		}
	}

	reset := func(arr [][]pair) {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr[i]); j++ {
				arr[i][j] = pair{-inf, -1}
			}
		}
	}

	update := func(u, v int, arr []pair) {
		if u >= arr[0].first {
			arr[1] = arr[0]
			arr[0] = pair{u, v}
		} else if u >= arr[1].first {
			arr[1] = pair{u, v}
		}
	}

	for i := 0; i < n; i++ {
		var cur = -inf
		for j := 1; j < m; j++ {
			copy(dp[i][j][0], col[j-1])
		}
		for j := 0; j < m; j++ {
			cur = max(cur, board[i][j])
			update(cur, i, col[j])
		}
	}

	reset(col)

	for i := 0; i < n; i++ {
		for j := m - 2; j >= 0; j-- {
			copy(dp[i][j][1], col[j+1])
		}
		var cur = -inf
		for j := m - 1; j >= 0; j-- {
			cur = max(cur, board[i][j])
			update(cur, i, col[j])
		}
	}

	reset(col)

	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			copy(dp[i][j][3], col[j-1])
		}
		var cur = -inf
		for j := 0; j < m; j++ {
			cur = max(cur, board[i][j])
			update(cur, i, col[j])
		}
	}

	reset(col)

	for i := n - 1; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			copy(dp[i][j][2], col[j+1])
		}
		var cur = -inf
		for j := m - 1; j >= 0; j-- {
			cur = max(cur, board[i][j])
			update(cur, i, col[j])
		}
	}

	check := func(a []pair, b []pair) int {
		var res = -inf

		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				// 不在同一行
				if a[i].second != b[j].second {
					res = max(res, a[i].first+b[j].first)
				}
			}
		}

		return res
	}

	var best = -inf

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := board[i][j] + check(dp[i][j][0], dp[i][j][1])
			best = max(best, tmp)
			tmp = board[i][j] + check(dp[i][j][0], dp[i][j][2])
			best = max(best, tmp)
			tmp = board[i][j] + check(dp[i][j][3], dp[i][j][1])
			best = max(best, tmp)
			tmp = board[i][j] + check(dp[i][j][3], dp[i][j][2])
			best = max(best, tmp)
		}
	}

	return int64(best)
}
