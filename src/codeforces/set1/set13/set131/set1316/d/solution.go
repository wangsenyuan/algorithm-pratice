package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(reader, 2*n)
	}

	res := solve(n, grid)

	if len(res) == 0 {
		fmt.Println("INVALID")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("VALID\n")
	for i := 0; i < n; i++ {
		buf.WriteString(res[i])
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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
var D = "DLUR"
var R = "URDL"

func solve(n int, grid [][]int) []string {
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			buf[i][j] = ' '
		}
	}

	que := make([]int, n*n)
	var head, tail int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := grid[i][2*j], grid[i][2*j+1]
			if x > 0 && y > 0 {
				x--
				y--
				if x == i && y == j {
					buf[x][y] = 'X'
					que[head] = x*n + y
					head++
				}
			}
		}
	}

	for tail < head {
		u, v := que[tail]/n, que[tail]%n
		tail++
		for i := 0; i < 4; i++ {
			r, c := u+dd[i], v+dd[i+1]
			if r >= 0 && r < n && c >= 0 && c < n &&
				buf[r][c] == ' ' &&
				grid[u][2*v] == grid[r][2*c] &&
				grid[u][2*v+1] == grid[r][2*c+1] {
				buf[r][c] = D[i]
				que[head] = r*n + c
				head++
			}
		}
	}

	bfs := func(x int, y int) {
		head, tail := 0, 0
		que[head] = x*n + y
		head++
		for tail < head {
			u, v := que[tail]/n, que[tail]%n
			tail++
			for i := 0; i < 4; i++ {
				r, c := u+dd[i], v+dd[i+1]
				if r >= 0 && r < n && c >= 0 && c < n &&
					buf[r][c] == ' ' &&
					grid[u][2*v] == grid[r][2*c] &&
					grid[u][2*v+1] == grid[r][2*c+1] {
					buf[r][c] = D[i]
					que[head] = r*n + c
					head++
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := grid[i][j*2], grid[i][j*2+1]

			if x > 0 && y > 0 && buf[i][j] == ' ' {
				// 这个格子没有被X访问到
				return nil
			}

			if x < 0 && y < 0 && buf[i][j] == ' ' {
				// ended in loop
				for k := 0; k < 4 && buf[i][j] == ' '; k++ {
					r, c := i+dd[k], j+dd[k+1]
					if r < 0 || r == n || c < 0 || c == n || buf[r][c] != ' ' || grid[r][2*c] > 0 {
						continue
					}
					// (-1, -1)
					// make them move to each other
					buf[i][j] = R[k]
					buf[r][c] = D[k]
					bfs(i, j)
					bfs(r, c)
				}
			}
			if buf[i][j] == ' ' {
				return nil
			}
		}
	}

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = string(buf[i])
	}

	return ans
}
