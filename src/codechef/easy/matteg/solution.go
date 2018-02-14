package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	T := readOneNum(scanner)

	for t := 0; t < T; t++ {
		firstLine := readNNums(scanner, 3)
		R, C, N := firstLine[0], firstLine[1], firstLine[2]
		secondLine := readNNums(scanner, 2)
		SX, SY := secondLine[0], secondLine[1]
		X := readNNums(scanner, N)
		Y := readNNums(scanner, N)
		board := make([][]int, R)
		for i := 0; i < R; i++ {
			board[i] = readNNums(scanner, C)
		}
		fmt.Println(solve(R, C, board, SX, SY, N, X, Y))
	}
}

func solve(R, C int, board [][]int, SX, SY int, N int, X []int, Y []int) int64 {
	NN := 1 << uint(N)
	cache := make([][]map[int]int64, R)
	for r := 0; r < R; r++ {
		cache[r] = make([]map[int]int64, C)
		for c := 0; c < C; c++ {
			cache[r][c] = make(map[int]int64)
		}
	}
	var dfs func(r, c int, state int) int64
	dfs = func(r, c int, state int) int64 {
		if state == NN-1 {
			return int64(board[r][c])
		}
		if v, found := cache[r][c][state]; found {
			return v
		}
		var tmp int64
		for i := 0; i < N; i++ {
			dx, dy := X[i], Y[i]
			if state&(1<<uint(i)) == 0 {
				if r-dx >= 0 && c-dy >= 0 {
					tmp = max(tmp, dfs(r-dx, c-dy, state|(1<<uint(i))))
				}
				if r-dx >= 0 && c+dy < C {
					tmp = max(tmp, dfs(r-dx, c+dy, state|(1<<uint(i))))
				}
				if r+dx < R && c-dy >= 0 {
					tmp = max(tmp, dfs(r+dx, c-dy, state|(1<<uint(i))))
				}
				if r+dx < R && c+dy < C {
					tmp = max(tmp, dfs(r+dx, c+dy, state|(1<<uint(i))))
				}
			}
		}
		cache[r][c][state] = tmp + int64(board[r][c])

		return cache[r][c][state]
	}

	return dfs(SX, SY, 0)
}

func solve1(R, C int, board [][]int, SX, SY int, N int, X []int, Y []int) int64 {
	NN := 1 << uint(N)
	// dp[r][c][state] = max value ending at (r, c), taking state paths
	dp := make([][][]int64, R)
	for r := 0; r < R; r++ {
		dp[r] = make([][]int64, C)
		for c := 0; c < C; c++ {
			dp[r][c] = make([]int64, NN)
		}
	}

	dp[SX][SY][0] = int64(board[SX][SY])

	for state := 0; state < NN; state++ {
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				if dp[r][c][state] > 0 {
					for i := 0; i < N; i++ {
						dx, dy := X[i], Y[i]
						if state&(1<<uint(i)) == 0 {
							if r-dx >= 0 && c-dy >= 0 {
								dp[r-dx][c-dy][state|(1<<uint(i))] =
									max(dp[r-dx][c-dy][state|(1<<uint(i))], dp[r][c][state]+int64(board[r-dx][c-dy]))
							}
							if r-dx >= 0 && c+dy < C {
								dp[r-dx][c+dy][state|(1<<uint(i))] =
									max(dp[r-dx][c+dy][state|(1<<uint(i))], dp[r][c][state]+int64(board[r-dx][c+dy]))
							}
							if r+dx < R && c-dy >= 0 {
								dp[r+dx][c-dy][state|(1<<uint(i))] =
									max(dp[r+dx][c-dy][state|(1<<uint(i))], dp[r][c][state]+int64(board[r+dx][c-dy]))
							}
							if r+dx < R && c+dy < C {
								dp[r+dx][c+dy][state|(1<<uint(i))] =
									max(dp[r+dx][c+dy][state|(1<<uint(i))], dp[r][c][state]+int64(board[r+dx][c+dy]))
							}
						}
					}
				}
			}
		}
	}

	var ans int64
	for state := 0; state < NN; state++ {
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				if dp[r][c][state] > ans {
					ans = dp[r][c][state]
				}
			}
		}
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
