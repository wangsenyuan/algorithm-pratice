package p1959

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	// 在8个方向上检查
	valid := func(x, y int) bool {
		return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
	}

	getColor := func(c byte) int {
		if c == 'B' {
			return -1
		}
		if c == 'W' {
			return 1
		}
		return 0
	}

	cur := getColor(color)

	check := func(dx, dy int) bool {
		x, y := rMove+dx, cMove+dy
		var cnt int

		for valid(x, y) && getColor(board[x][y]) == -cur {
			x += dx
			y += dy
			cnt++
		}
		// out of board
		if !valid(x, y) {
			return false
		}
		if cnt == 0 || getColor(board[x][y]) != cur {
			return false
		}
		return true
	}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if check(dx, dy) {
				return true
			}
		}
	}
	return false
}
