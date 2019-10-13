package p1222

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var res [][]int

	x, y := king[0], king[1]

	board := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		board[i] = make([]bool, 8)
	}

	for _, queen := range queens {
		a, b := queen[0], queen[1]
		board[a][b] = true
	}

	for i := x - 1; i >= 0; i-- {
		if board[i][y] {
			res = append(res, []int{i, y})
			break
		}
	}

	for i := x + 1; i < 8; i++ {
		if board[i][y] {
			res = append(res, []int{i, y})
			break
		}
	}

	for j := y - 1; j >= 0; j-- {
		if board[x][j] {
			res = append(res, []int{x, j})
			break
		}
	}
	for j := y + 1; j < 8; j++ {
		if board[x][j] {
			res = append(res, []int{x, j})
			break
		}
	}

	for i := 1; x-i >= 0 && y-i >= 0; i++ {
		if board[x-i][y-i] {
			res = append(res, []int{x - i, y - i})
			break
		}
	}

	for i := 1; x+i < 8 && y+i < 8; i++ {
		if board[x+i][y+i] {
			res = append(res, []int{x + i, y + i})
			break
		}
	}

	for i := 1; x-i >= 0 && y+i < 8; i++ {
		if board[x-i][y+i] {
			res = append(res, []int{x - i, y + i})
			break
		}
	}

	for i := 1; x+i < 8 && y-i >= 0; i++ {
		if board[x+i][y-i] {
			res = append(res, []int{x + i, y - i})
			break
		}
	}

	return res
}
