package p794

func validTicTacToe(board []string) bool {
	x, o := 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				x++
			} else if board[i][j] == 'O' {
				o++
			}
		}
	}

	if x-o > 1 || o > x {
		return false
	}

	if board[1][1] != ' ' && board[0][0] == board[1][1] && board[2][2] == board[1][1] && board[0][2] == board[1][1] && board[2][0] == board[1][1] {
		return false
	}

	winTimes := func(t byte) (cnt int) {
		for i := 0; i < 3; i++ {
			if board[i][0] == t && board[i][1] == t && board[i][2] == t {
				cnt++
			}
		}
		for i := 0; i < 3; i++ {
			if board[0][i] == t && board[1][i] == t && board[2][i] == t {
				cnt++
			}
		}
		if board[1][1] == t && board[0][0] == board[1][1] && board[2][2] == board[1][1] {
			cnt++
		}
		if board[1][1] == t && board[0][2] == board[1][1] && board[2][0] == board[1][1] {
			cnt++
		}
		return cnt
	}

	xw := winTimes('X')
	ow := winTimes('O')

	if xw > 0 && ow > 0 {
		return false
	}

	if xw > 1 || ow > 1 {
		return false
	}

	if xw == 1 && (o+x)%2 == 0 {
		return false
	}

	if ow == 1 && (o+x)%2 == 1 {
		return false
	}

	return true
}
