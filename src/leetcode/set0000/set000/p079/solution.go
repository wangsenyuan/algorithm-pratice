package p079

var dd = []int{-1, 0, 1, 0, -1}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	checked := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		checked[i] = make([]bool, len(board[i]))
	}

	var find func(i, j int, word string) bool

	find = func(i, j int, word string) bool {
		checked[i][j] = true

		for k := 0; k < 4; k++ {
			x := i + dd[k]
			y := j + dd[k+1]
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) {
				if !checked[x][y] {
					if board[x][y] == word[0] {
						if len(word) == 1 || find(x, y, word[1:]) {
							return true
						}
					}
				}
			}
		}

		checked[i][j] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if len(word) == 1 || find(i, j, word[1:]) {
					return true
				}
			}
		}
	}

	return false
}
