package p289

func gameOfLife(board [][]int) {
	// 1 + 8 = die
	// board[i][j] stay same = live
	// 0 + 8 = relive

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			var tmp int
			if i > 0 && j > 0 {
				tmp += board[i-1][j-1] % 8
			}
			if i > 0 {
				tmp += board[i-1][j] % 8
			}
			if i > 0 && j < len(board[i])-1 {
				tmp += board[i-1][j+1] % 8
			}
			if j > 0 {
				tmp += board[i][j-1] % 8
			}
			if j < len(board[i])-1 {
				tmp += board[i][j+1] % 8
			}
			if i < len(board)-1 && j > 0 {
				tmp += board[i+1][j-1] % 8
			}
			if i < len(board)-1 {
				tmp += board[i+1][j] % 8
			}
			if i < len(board)-1 && j < len(board[i])-1 {
				tmp += board[i+1][j+1] % 8
			}
			if board[i][j] == 1 {
				if tmp < 2 || tmp > 3 {
					board[i][j] = 9
				}
			} else {
				if tmp == 3 {
					board[i][j] = 8
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 8 {
				board[i][j] = 1
			} else if board[i][j] == 9 {
				board[i][j] = 0
			}
		}
	}

}
