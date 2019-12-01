package p1275

func tictactoe(moves [][]int) string {
	var res int

	grid := make([]byte, 9)

	put := func(c byte, move []int) {
		x, y := move[0], move[1]
		grid[x*3+y] = c
	}

	checkWin := func(c byte) bool {
		for i := 0; i < 3; i++ {
			var j int
			for j < 3 && grid[i*3+j] == c {
				j++
			}
			if j == 3 {
				return true
			}
		}

		for j := 0; j < 3; j++ {
			var i int
			for i < 3 && grid[i*3+j] == c {
				i++
			}
			if i == 3 {
				return true
			}
		}

		if grid[0] == c && grid[4] == c && grid[8] == c {
			return true
		}

		if grid[2] == c && grid[4] == c && grid[6] == c {
			return true
		}
		return false
	}

	checkWinner := func() int {
		if checkWin('X') {
			return 1
		}
		if checkWin('O') {
			return 2
		}
		return 0
	}

	for i := 0; i < len(moves) && res == 0; i++ {
		if i%2 == 0 {
			put('X', moves[i])
		} else {
			put('O', moves[i])
		}
		res = checkWinner()
	}

	if res == 0 {
		if len(moves) == 9 {
			return "Draw"
		}
		return "Pending"
	}
	if res == 1 {
		return "A"
	}
	return "B"
}
