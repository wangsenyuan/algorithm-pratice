package main

import "fmt"

func main() {
	input := []string{"..9748...", "7........", ".2.1.9...", "..7...24.", ".64.1.59.", ".98...3..", "...8.3.2.", "........6", "...2759.."}
	board := make([][]byte, 9, 9)
	for i := range board {
		board[i] = []byte(input[i])
	}
	solveSudoku(board)
	for i := range board {
		fmt.Printf("%s\n", string(board[i]))
	}
}

func solveSudoku(board [][]byte) {
	var row = make([]int, 9, 9)
	var col = make([]int, 9, 9)
	var grid = make([][]int, 3, 3)

	for i := range grid {
		grid[i] = make([]int, 3, 3)
	}

	fill := func(i, j int, x uint) {
		row[i] |= 1 << x
		col[j] |= 1 << x
		grid[i/3][j/3] |= 1 << x
	}

	erase := func(i, j int, x uint) {
		row[i] &= ^(1 << x)
		col[j] &= ^(1 << x)
		grid[i/3][j/3] &= ^(1 << x)
	}

	var dfs func(i, j int)

	dfs = func(i, j int) {
		if i == 9 {
			panic("solution found")
		}
		a := i
		b := j

		if j == 8 {
			a++
			b = 0
		} else {
			b++
		}

		if board[i][j] != '.' {
			dfs(a, b)
			return
		}

		for x := uint(1); x <= 9; x++ {
			if (row[i]&(1<<x)) > 0 || (col[j]&(1<<x)) > 0 || (grid[i/3][j/3]&(1<<x)) > 0 {
				continue
			}

			board[i][j] = byte('0' + int(x))
			fill(i, j, x)
			dfs(a, b)
			erase(i, j, x)
			board[i][j] = '.'
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			x := board[i][j] - '0'
			fill(i, j, uint(x))
		}
	}

	defer func() {
		if r := recover(); r != nil {

		}
	}()
	dfs(0, 0)
}
