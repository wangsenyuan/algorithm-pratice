package p2018

func placeWordInCrossword(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	k := len(word)

	checkRow := func(r int, a int, b int) bool {
		l := b - a - 1
		if l != k {
			return false
		}
		var j int
		for i := a + 1; i < b; i++ {
			if board[r][i] == ' ' || board[r][i] == word[j] {
				j++
				continue
			}
			break
		}

		if j == k {
			return true
		}

		j = 0
		for i := b - 1; i > a; i-- {
			if board[r][i] == ' ' || board[r][i] == word[j] {
				j++
				continue
			}
			return false
		}

		return true
	}

	for r := 0; r < m; r++ {
		cur := -1
		for c := 0; c <= n; c++ {
			if c == n || board[r][c] == '#' {
				if checkRow(r, cur, c) {
					return true
				}
				cur = c
				//a candidate
			}
		}
	}

	checkCol := func(c int, a int, b int) bool {
		l := b - a - 1
		if l != k {
			return false
		}
		var j int
		for i := a + 1; i < b; i++ {
			if board[i][c] == ' ' || board[i][c] == word[j] {
				j++
				continue
			}
			break
		}
		if j == k {
			return true
		}

		j = 0
		for i := b - 1; i > a; i-- {
			if board[i][c] == ' ' || board[i][c] == word[j] {
				j++
				continue
			}
			return false
		}

		return true
	}

	for c := 0; c < n; c++ {
		cur := -1
		for r := 0; r <= m; r++ {
			if r == m || board[r][c] == '#' {
				if checkCol(c, cur, r) {
					return true
				}
				cur = r
			}
		}
	}

	return false
}
