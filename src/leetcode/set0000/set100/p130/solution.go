package p130

const (
	X = 'X'
	O = 'O'
	Y = 'Y'
)

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		flip(board, i, 0)
		flip(board, i, len(board[i])-1)
	}

	for j := 0; j < len(board[0]); j++ {
		flip(board, 0, j)
		flip(board, len(board)-1, j)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == Y {
				board[i][j] = O
			} else {
				board[i][j] = X
			}
		}
	}
}

func flip(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != O {
		return
	}
	board[i][j] = Y
	flip(board, i, j+1)
	flip(board, i+1, j)
	flip(board, i-1, j)
	flip(board, i, j-1)
}
