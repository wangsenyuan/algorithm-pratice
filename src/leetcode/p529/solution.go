package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'M', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
	}
	click := []int{3, 0}
	result := updateBoard(board, click)

	for _, row := range result {
		for _, b := range row {
			fmt.Printf("%c ", b)
		}
		fmt.Println()
	}
}

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func updateBoard(board [][]byte, click []int) [][]byte {

	valid := func(x, y int) bool {
		return x >= 0 && x < len(board) && y >= 0 && y < len(board[x])
	}

	var reveal func(i, j int)

	reveal = func(i, j int) {
		cnt := 0

		for k := 0; k < 8; k++ {
			x, y := i+dx[k], j+dy[k]
			if !valid(x, y) {
				continue
			}
			if board[x][y] == 'M' {
				cnt += 1
			}
		}

		if cnt > 0 {
			board[i][j] = byte('0' + cnt)
			return
		}

		board[i][j] = 'B'

		for k := 0; k < 8; k++ {
			x, y := i+dx[k], j+dy[k]
			if !valid(x, y) {
				continue
			}

			if board[x][y] == 'E' {
				reveal(x, y)
			}
		}
	}

	x, y := click[0], click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	reveal(x, y)

	return board
}
