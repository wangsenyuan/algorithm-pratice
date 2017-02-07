package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n int
		fmt.Scanf("%d", &n)
		board := make([][]byte, n)

		for i := 0; i < n; i++ {
			var s string
			fmt.Scanf("%s", &s)
			board[i] = []byte(s)
		}

		res := capturePawn(board, n)

		fmt.Println(res)

		t--
	}
}

func capturePawn(board [][]byte, n int) int {
	dp := make([][]int, n)

	ki, kj := 0, 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if board[i][j] == 'P' {
				dp[i][j] = 1
			} else if board[i][j] == 'K' {
				ki = i
				kj = j
			}
		}
	}

	dpv := func(i, j int) int {
		if i < 0 || j < 0 || i >= n || j >= n {
			return 0
		} else {
			return dp[i][j]
		}
	}
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < n; i++ {
			dp[i][j] += max4(dpv(i-2, j+1), dpv(i+2, j+1), dpv(i-1, j+2), dpv(i+1, j+2))
		}
	}

	return dp[ki][kj]
}

func max4(a, b, c, d int) int {
	x := a
	if x < b {
		x = b
	}

	if x < c {
		x = c
	}
	if x < d {
		x = d
	}
	return x
}
