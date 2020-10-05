package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	board := make([]string, 3)
	for i := 0; i < 3; i++ {
		board[i], _ = reader.ReadString('\n')
	}
	fmt.Println(solve(board))
}

func solve(board []string) string {
	cnt := make([]int, 2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				cnt[0]++
			} else if board[i][j] == '0' {
				cnt[1]++
			}
		}
	}

	if cnt[1] > cnt[0] || cnt[0]-cnt[1] > 1 {
		return "illegal"
	}

	win := func(x byte) bool {
		for i := 0; i < 3; i++ {
			var ok = true
			for j := 0; j < 3 && ok; j++ {
				ok = board[i][j] == x
			}
			if ok {
				return true
			}
		}

		for j := 0; j < 3; j++ {
			var ok = true
			for i := 0; i < 3 && ok; i++ {
				ok = board[i][j] == x
			}
			if ok {
				return true
			}
		}
		if board[0][0] == x && board[1][1] == x && board[2][2] == x {
			return true
		}
		if board[0][2] == x && board[1][1] == x && board[2][0] == x {
			return true
		}
		return false
	}
	a := win('X')
	b := win('0')
	if a && b {
		return "illegal"
	}

	if b && cnt[0] > cnt[1] {
		return "illegal"
	}
	if a && cnt[0] == cnt[1] {
		return "illegal"
	}
	// determine win next
	if a {
		return "the first player won"
	}
	if b {
		return "the second player won"
	}
	if cnt[0]+cnt[1] == 9 {
		return "draw"
	}
	if cnt[0] > cnt[1] {
		return "second"
	}
	return "first"
}
