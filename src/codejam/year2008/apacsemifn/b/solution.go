package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	f, err := os.Open("./B-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		line := readNNums(scanner, 4)
		R, C := line[1], line[0]
		r, c := line[3], line[2]
		board := make([][]int, R)
		for i := 0; i < R; i++ {
			board[i] = readNNums(scanner, C)
		}
		res := solve(R, C, r, c, board)
		if res == INF {
			fmt.Printf("Case #%d: forever\n", i)
		} else {
			fmt.Printf("Case #%d: %d day(s)\n", i, res)
		}
	}
}

const INF = 1001

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}

func solve(R, C int, r, c int, board [][]int) int {

	var dfs func(days int)
	var best int

	r--
	c--

	dfs = func(days int) {
		if board[r][c] == 0 {
			return
		}

		safe := true
		for i := 0; i < 4; i++ {
			a, b := dx[i]+r, dy[i]+c
			if a >= 0 && a < R && b >= 0 && b < C && board[a][b] > 0 {
				safe = false
				break
			}
		}
		if safe {
			best = INF
			return
		}
		if best == INF {
			return
		}

		if days > best {
			best = days
		}

		prev := make([][]int, R)
		dmg := make([][]int, R)
		for i := 0; i < R; i++ {
			prev[i] = make([]int, C)
			copy(prev[i], board[i])
			dmg[i] = make([]int, C)
		}

		for i := 0; i < R; i++ {
			for j := 0; j < C; j++ {
				if i != r || j != c {
					u, v := -1, -1
					for k := 0; k < 4; k++ {
						a, b := dx[k]+i, dy[k]+j
						if a >= 0 && a < R && b >= 0 && b < C && (u < 0 || board[a][b] > board[u][v]) {
							u, v = a, b
						}
					}
					dmg[u][v] += board[i][j]
				}
			}
		}

		sp := board[r][c]

		for i := 0; i < R; i++ {
			for j := 0; j < C; j++ {
				prev[i][j] -= dmg[i][j]
			}
		}

		if r > 0 && prev[r-1][c] > 0 {
			for i := 0; i < R; i++ {
				for j := 0; j < C; j++ {
					board[i][j] = prev[i][j]
					if i == r-1 && j == c {
						board[i][j] -= sp
					}
					if board[i][j] < 0 {
						board[i][j] = 0
					}
				}
			}
			dfs(days + 1)
		}

		if r < R-1 && prev[r+1][c] > 0 {
			for i := 0; i < R; i++ {
				for j := 0; j < C; j++ {
					board[i][j] = prev[i][j]
					if i == r+1 && j == c {
						board[i][j] -= sp
					}
					if board[i][j] < 0 {
						board[i][j] = 0
					}
				}
			}
			dfs(days + 1)
		}

		if c > 0 && prev[r][c-1] > 0 {
			for i := 0; i < R; i++ {
				for j := 0; j < C; j++ {
					board[i][j] = prev[i][j]
					if i == r && j == c-1 {
						board[i][j] -= sp
					}
					if board[i][j] < 0 {
						board[i][j] = 0
					}
				}
			}
			dfs(days + 1)
		}

		if c < C-1 && prev[r][c+1] > 0 {
			for i := 0; i < R; i++ {
				for j := 0; j < C; j++ {
					board[i][j] = prev[i][j]
					if i == r && j == c+1 {
						board[i][j] -= sp
					}
					if board[i][j] < 0 {
						board[i][j] = 0
					}
				}
			}
			dfs(days + 1)
		}

		for i := 0; i < R; i++ {
			for j := 0; j < C; j++ {
				board[i][j] = prev[i][j]
				if board[i][j] < 0 {
					board[i][j] = 0
				}
			}
		}
		dfs(days + 1)
	}

	dfs(0)

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
