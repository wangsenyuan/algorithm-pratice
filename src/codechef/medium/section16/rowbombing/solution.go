package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)[:m]
		}
		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const INF = 1 << 30

var dd = []int{-1, 0, 1, 0, -1}

func solve(board []string) int {
	m := len(board)
	n := len(board[0])

	bombs := make([][][]int, m)
	for i := 0; i < m; i++ {
		bombs[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			bombs[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				bombs[i][j][k] = INF
			}
		}
	}

	que := make([]Item, m*n*4)
	// in the middle
	var front, end int = m * n * 2, m * n * 2

	push_front := func(x, y, k int, d int) {
		if bombs[x][y][k] <= d {
			return
		}
		bombs[x][y][k] = d
		front--
		que[front] = Item{x, y, k}
	}

	push_back := func(x, y, k int, d int) {
		if bombs[x][y][k] <= d {
			return
		}
		bombs[x][y][k] = d
		que[end] = Item{x, y, k}
		end++
	}

	push_back(0, 0, 0, 0)

	for front < end {
		cur := que[front]
		front++
		x, y, lv := cur.x, cur.y, cur.k
		d := bombs[x][y][lv]

		for i := 0; i < 4; i++ {
			u, v := x+dd[i], y+dd[i+1]
			if u >= 0 && u < m && v >= 0 && v < n {
				if x == u && (lv == 1 || board[u][v] == '.') {
					push_front(u, v, lv, d)
				} else if u != x && board[u][v] == '.' {
					push_front(u, v, 0, d)
				}
			}
		}
		for i := 0; i < 3; i++ {
			dx := dd[i]
			if x+dx < 0 || x+dx >= m {
				continue
			}
			push_back(x+dx, y, 1, d+1)
		}
	}

	return min(bombs[m-1][n-1][0], bombs[m-1][n-1][1])
}

type Item struct {
	x int
	y int
	k int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
