package main

import "fmt"

func candyCrush(board [][]int) [][]int {
	n := len(board)
	if n == 0 {
		return board
	}
	m := len(board[0])
	if m == 0 {
		return board
	}

	flag := make([][]bool, n)
	for i := 0; i < n; i++ {
		flag[i] = make([]bool, m)
	}

	check := func(x, y int) bool {
		a, b := x, x
		c, d := y, y

		for a > 0 && board[x][y] == board[a-1][y] {
			a--
		}

		for b < n-1 && board[x][y] == board[b+1][y] {
			b++
		}

		for c > 0 && board[x][y] == board[x][c-1] {
			c--
		}

		for d < m-1 && board[x][y] == board[x][d+1] {
			d++
		}

		if b-a+1 >= 3 {
			for i := a; i <= b; i++ {
				flag[i][y] = true
			}

		}

		if d-c+1 >= 3 {
			for j := c; j <= d; j++ {
				flag[x][j] = true
			}
		}

		return b-a >= 2 || d-c >= 2
	}

	canCrush := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] > 0 {
				canCrush = check(i, j) || canCrush
			}
		}
	}

	if !canCrush {
		return board
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if flag[i][j] {
				k := i - 1
				for k >= 0 && (board[k][j] == 0 || flag[k][j]) {
					k--
				}
				if k >= 0 {
					board[i][j] = board[k][j]
					flag[k][j] = true
				} else {
					board[i][j] = 0
				}
			}
		}
	}
	return candyCrush(board)
}

func main() {
	board := [][]int{{110, 5, 112, 113, 114},
		{210, 211, 5, 213, 214},
		{310, 311, 3, 313, 314},
		{410, 411, 412, 5, 414},
		{5, 1, 512, 3, 3},
		{610, 4, 1, 613, 614},
		{710, 1, 2, 713, 714},
		{810, 1, 2, 1, 1},
		{1, 1, 2, 2, 2},
		{4, 1, 4, 4, 1014},
	}
	result := candyCrush(board)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
