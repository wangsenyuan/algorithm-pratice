package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)[:m]
	}

	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 3)
	}

	res := solve(grid, queries)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
}

const inf = 1 << 30

var dd = []int{-1, 0, 1, 0, -1}

func solve(grid []string, queries [][]int) []int {
	ans := make([]int, len(queries))
	if checkChessboard(grid) {
		for i, cur := range queries {
			l, r := cur[0]-1, cur[1]-1
			ans[i] = int(grid[l][r] - '0')
		}
		return ans
	}
	n := len(grid)
	m := len(grid[0])

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		dist[r][c] = 0
		for i := 0; i < 4; i++ {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == grid[r][c] && dist[r][c] < 0 {
				dfs(x, y)
			}
		}
	}

	check := func(r int, c int) bool {
		if r > 0 && grid[r-1][c] == grid[r][c] {
			return true
		}
		if r+1 < n && grid[r+1][c] == grid[r][c] {
			return true
		}
		if c > 0 && grid[r][c-1] == grid[r][c] {
			return true
		}
		if c+1 < m && grid[r][c+1] == grid[r][c] {
			return true
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dist[i][j] < 0 && check(i, j) {
				dfs(i, j)
			}
		}
	}

	// 还有一部分没有被处理的，它们都是孤立在周围的
	que := make([]int, n*m)
	var head, tail int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dist[i][j] == 0 {
				que[head] = i*m + j
				head++
			}
		}
	}

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++
		for i := 0; i < 4; i++ {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && dist[x][y] < 0 {
				// 在下个周期，(r, c)会变成和(x, y)同样的颜色，然后(x, y)和(r, c)共频
				dist[x][y] = dist[r][c] + 1
				que[head] = x*m + y
				head++
			}
		}
	}

	for i, cur := range queries {
		l, r, p := cur[0]-1, cur[1]-1, cur[2]
		ans[i] = int(grid[l][r] - '0')

		if p < dist[l][r] {
			continue
		}
		diff := p - dist[l][r]
		diff %= 2
		if diff == 1 {
			ans[i] ^= 1
		}
	}

	return ans
}

func checkChessboard(grid []string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dx := (i + j) % 2
			if dx == 0 && grid[i][j] != grid[0][0] {
				return false
			}
			if dx == 1 && grid[i][j] == grid[0][0] {
				return false
			}
		}
	}
	return true
}
