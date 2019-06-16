package p999

func numRookCaptures(board [][]byte) int {
	if len(board) == 0 {
		return 0
	}
	if len(board[0]) == 0 {
		return 0
	}

	m := len(board)
	n := len(board[0])

	findRook := func() (x int, y int) {
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if board[i][j] == 'R' {
					x = i
					y = j
					return
				}
			}
		}
		return
	}

	x, y := findRook()

	if board[x][y] != 'R' {
		return 0
	}

	var ans int

	u := x - 1
	for u >= 0 && board[u][y] == '.' {
		u--
	}
	if u >= 0 && board[u][y] == 'p' {
		ans++
	}
	u = x + 1
	for u < m && board[u][y] == '.' {
		u++
	}
	if u < m && board[u][y] == 'p' {
		ans++
	}
	v := y - 1
	for v >= 0 && board[x][v] == '.' {
		v--
	}
	if v >= 0 && board[x][v] == 'p' {
		ans++
	}

	v = y + 1
	for v < n && board[x][v] == '.' {
		v++
	}
	if v < n && board[x][v] == 'p' {
		ans++
	}
	return ans
}
