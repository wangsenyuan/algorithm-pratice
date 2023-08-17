package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = readString(reader)[:m]
	}
	res := solve(board, k)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(board []string, k int) int {
	points := findValidPoints(board)
	n := len(points)
	// 0 for start, n - 1 for dest

	h := len(board)
	w := len(board[0])

	dist := make([][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dist[i][j] = -1
		}
	}

	que := make([]int, h*w)
	var front, end int
	pushBack := func(x int, y int) {
		que[end] = x*w + y
		end++
	}

	popFront := func() (x int, y int) {
		cur := que[front]
		x, y = cur/w, cur%w
		front++
		return
	}

	reset := func() {
		front = 0
		end = 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				dist[i][j] = -1
			}
		}
	}

	bfs := func(x int, y int) {
		reset()
		dist[x][y] = 0
		pushBack(x, y)

		for front < end {
			a, b := popFront()
			for i := 0; i < 4; i++ {
				c, d := a+dd[i], b+dd[i+1]
				if c >= 0 && c < h && d >= 0 && d < w && dist[c][d] < 0 && board[c][d] != '#' {
					dist[c][d] = dist[a][b] + 1
					pushBack(c, d)
				}
			}
		}
	}

	p_dist := make([][]int, n)
	for i := 0; i < n; i++ {
		p_dist[i] = make([]int, n)
		bfs(points[i][0], points[i][1])
		for j := 0; j < n; j++ {
			p_dist[i][j] = dist[points[j][0]][points[j][1]]
		}
	}

	N := 1 << n
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// dp[i][j] 表示从0出发, 且访问往i表示的节点后，停在节点j时的最大距离
			dp[i][j] = -(k + 10)
		}
	}

	dp[1][0] = 0

	for state := 1; state < N; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 0 || dp[state][i] < 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if (state>>j)&1 == 0 && p_dist[i][j] > 0 && dp[state][i]+p_dist[i][j] <= k {
					// can reach
					if dp[state|(1<<j)][j] < 0 || dp[state|(1<<j)][j] >= dp[state][i]+p_dist[i][j] {
						dp[state|(1<<j)][j] = dp[state][i] + p_dist[i][j]
					}
				}
			}
		}
	}

	best := -1

	for state := 1; state < N; state++ {
		if dp[state][n-1] > 0 {
			dc := digitCount(state)
			best = max(best, dc-2)
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func findValidPoints(board []string) [][]int {
	var cnt int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'S' {
				cnt++
			} else if board[i][j] == 'G' {
				cnt++
			} else if board[i][j] == 'o' {
				cnt++
			}
		}
	}
	points := make([][]int, cnt)
	var it = 1
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'S' {
				points[0] = []int{i, j}
			} else if board[i][j] == 'G' {
				points[cnt-1] = []int{i, j}
			} else if board[i][j] == 'o' {
				points[it] = []int{i, j}
				it++
			}
		}
	}
	return points
}
