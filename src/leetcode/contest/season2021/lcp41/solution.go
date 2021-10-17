package lcp41

func flipChess(chessboard []string) int {
	// X for black, O for white
	m := len(chessboard)
	n := len(chessboard[0])

	buf := make([][]byte, m)
	for i := 0; i < m; i++ {
		buf[i] = []byte(chessboard[i])
	}

	check := func(x, y, dx, dy int, flip bool) bool {
		for x >= 0 && x < m && y >= 0 && y < n && buf[x][y] == 'O' {
			if flip {
				buf[x][y] = 'X'
			}
			x += dx
			y += dy
		}
		return x >= 0 && x < m && y >= 0 && y < n && buf[x][y] == 'X'
	}

	var flip func()

	flip = func() {
		var flag bool
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if buf[i][j] == 'O' {
					tmp := false
					if check(i, j, -1, -1, false) && check(i, j, 1, 1, false) {
						tmp = true
						check(i-1, j-1, -1, -1, true)
						check(i+1, j+1, 1, 1, true)
					}
					if check(i, j, -1, 0, false) && check(i, j, 1, 0, false) {
						tmp = true
						check(i-1, j, -1, 0, true)
						check(i+1, j, 1, 0, true)
					}
					if check(i, j, -1, 1, false) && check(i, j, 1, -1, false) {
						tmp = true
						check(i-1, j+1, -1, 1, true)
						check(i+1, j-1, 1, -1, true)
					}
					if check(i, j, 0, -1, false) && check(i, j, 0, 1, false) {
						tmp = true
						check(i, j-1, 0, -1, true)
						check(i, j+1, 0, 1, true)
					}
					if tmp {
						buf[i][j] = 'X'
						flag = true
					}
				}
			}
		}
		if flag {
			flip()
		}
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if chessboard[i][j] == '.' {
				for r := 0; r < m; r++ {
					buf[r] = []byte(chessboard[r])
				}
				buf[i][j] = 'X'
				flip()
				var cnt int
				for r := 0; r < m; r++ {
					for c := 0; c < n; c++ {
						if buf[r][c] != chessboard[r][c] {
							cnt++
						}
					}
				}
				cnt--
				if cnt > ans {
					ans = cnt
				}
			}
		}
	}
	return ans
}
