package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte("X..X"),
		[]byte("...X"),
		[]byte("...X"),
	}

	fmt.Println(countBattleships(board))
}

func countBattleships(board [][]byte) int {
	m := len(board)
	if m == 0 {
		return 0
	}
	n := len(board[0])
	if n == 0 {
		return 0
	}

	label := make([][]int, m)
	idx := 0

	for i := 0; i < m; i++ {
		label[i] = make([]int, n)
		for j := 0; j < n; j++ {
			label[i][j] = -1

			if board[i][j] == '.' {
				continue
			}

			if i > 0 && board[i-1][j] == 'X' {
				label[i][j] = label[i-1][j]
				continue
			}

			if j > 0 && board[i][j-1] == 'X' {
				label[i][j] = label[i][j-1]
				continue
			}

			label[i][j] = idx
			idx++
		}
	}

	return idx
}
