package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		readString(reader)
		n, m := readTwoNums(reader)
		n++
		m++
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)[:m]
		}
		x := readNum(reader)
		res := solve(grid, x)
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(grid []string, x int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, x+1)
			for k := 0; k <= x; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	X := x + 1
	que := make([]int, n*m*X)
	var front, end int

	push := func(i, j, k int) {
		que[end] = i*m*X + j*X + k
		end++
	}

	pop := func() (i int, j int, k int) {
		cur := que[front]
		i = cur / (m * X)
		j = (cur % (m * X)) / X
		k = cur % X
		front++
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' {
				push(i, j, 0)
				dp[i][j][0] = 0
			}
		}
	}

	for front < end {
		i, j, k := pop()
		if k == x {
			return dp[i][j][k]
		}
		for d := 0; d < 4; d++ {
			u, v := i+dd[d], j+dd[d+1]
			if u >= 0 && u < n && v >= 0 && v < m && grid[u][v] != '#' {
				// not blocked
				kk := k
				if grid[u][v] >= '0' && grid[u][v] <= '9' {
					kk += int(grid[u][v] - '0')
				}
				if kk <= x && dp[u][v][kk] < 0 {
					dp[u][v][kk] = dp[i][j][k] + 1
					push(u, v, kk)
				}
			}
		}
	}

	res := -1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dp[i][j][x] > 0 {
				if res < 0 || res > dp[i][j][x] {
					res = dp[i][j][x]
				}
			}
		}
	}
	return res
}
