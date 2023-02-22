package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	m, n := readTwoNums(reader)
	board := make([]string, m)
	for i := 0; i < m; i++ {
		board[i] = readString(reader)[:n]
	}
	res := solve(board)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(board []string) int {
	m := len(board)
	n := len(board[0])
	dist := make([][]int, m)
	check_pos := []int{-1, -1}
	spoon_pos := []int{-1, -1}
	que := make([]int, m*n)
	var front, end int
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
			if board[i][j] == 'D' {
				dist[i][j] = 0
				que[end] = i*n + j
				end++
			} else if board[i][j] == '@' {
				check_pos = []int{i, j}
			} else if board[i][j] == '$' {
				spoon_pos = []int{i, j}
			}

		}
	}

	for front < end {
		cur := que[front]
		front++
		x, y := cur/n, cur%n
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				u, v := x+i, y+j
				if u >= 0 && u < m && v >= 0 && v < n {
					if dist[u][v] < 0 {
						dist[u][v] = dist[x][y] + 1
						que[end] = u*n + v
						end++
					}
				}
			}
		}
	}
	vis := make([]bool, m*n)
	// can chef pass through cells have dist at least x apart from monsters
	check := func(expect int) bool {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				vis[i*n+j] = false
			}
		}
		front = 0
		end = 0
		que[end] = check_pos[0]*n + check_pos[1]
		vis[que[0]] = true
		end++

		for front < end {
			cur := que[front]
			front++
			if cur == spoon_pos[0]*n+spoon_pos[1] {
				return true
			}
			x, y := cur/n, cur%n
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					u, v := x+i, y+j
					if u >= 0 && u < m && v >= 0 && v < n && dist[u][v] >= expect && !vis[u*n+v] {
						vis[u*n+v] = true
						que[end] = u*n + v
						end++
					}
				}
			}
		}
		return false
	}

	l, r := 1, min(dist[check_pos[0]][check_pos[1]], dist[spoon_pos[0]][spoon_pos[1]])+1

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
