package p036

func isValidSudoku(board [][]byte) bool {

	rows := make([]int, 9)
	cols := make([]int, 9)
	quart := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			x := uint(board[i][j] - '0')
			if rows[i]&(1<<x) > 0 {
				return false
			}
			rows[i] |= 1 << x
			if cols[j]&(1<<x) > 0 {
				return false
			}
			cols[j] |= 1 << x

			k := (i/3)*3 + (j / 3)
			if quart[k]&(1<<x) > 0 {
				return false
			}
			quart[k] |= 1 << x
		}
	}

	return true
}
