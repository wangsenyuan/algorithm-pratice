package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		color := make([]string, n)
		for i := 0; i < n; i++ {
			color[i] = readString(reader)[:m]
		}
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)[:m]
		}
		res := solve(color, grid)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(color []string, grid []string) []int {
	n := len(grid)
	m := len(grid[0])

	move := func(u, v int) (int, int) {
		if u < 0 || u >= n || v < 0 || v >= m {
			return -1, -1
		}
		if grid[u][v] == 'U' {
			u--
		} else if grid[u][v] == 'R' {
			v++
		} else if grid[u][v] == 'D' {
			u++
		} else {
			v--
		}
		return u, v
	}

	h := bits.Len(uint(n*m + 1))

	dp := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r, c := move(i, j)
			dp[i*m+j] = r*m + c
		}
	}
	ndp := make([]int, n*m)

	for d := 0; d <= h; d++ {
		for i := 0; i < n*m; i++ {
			ndp[i] = dp[dp[i]]
		}
		copy(dp, ndp)
	}

	tot := make(map[int]bool)
	blk := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			id := dp[i*m+j]
			tot[id] = true
			if color[i][j] == '0' {
				blk[id] = true
			}
		}
	}

	return []int{len(tot), len(blk)}
}

func solve1(color []string, grid []string) []int {
	n := len(grid)
	m := len(grid[0])

	move := func(u, v int) (int, int) {
		if u < 0 || u >= n || v < 0 || v >= m {
			return -1, -1
		}
		if grid[u][v] == 'U' {
			u--
		} else if grid[u][v] == 'R' {
			v++
		} else if grid[u][v] == 'D' {
			u++
		} else {
			v--
		}
		return u, v
	}

	deg := make([][]int, n)
	for i := 0; i < n; i++ {
		deg[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				r, c := i+dd[k], j+dd[k+1]
				x, y := move(r, c)
				if x == i && y == j {
					deg[i][j]++
				}
			}
		}
	}

	que := make([]int, n*m)

	vis := make([][]bool, n)

	var head, tail int
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			if deg[i][j] == 0 {
				vis[i][j] = true
				que[head] = i*m + j
				head++
			}
		}
	}

	for tail < head {
		u, v := que[tail]/m, que[tail]%m
		tail++
		r, c := move(u, v)
		if r >= 0 && r < n && c >= 0 && c < m {
			deg[r][c]--
			if deg[r][c] == 0 && !vis[r][c] {
				vis[r][c] = true
				que[head] = r*m + c
				head++
			}
		}

	}

	res := []int{0, 0}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	ok := make([]bool, n*m+1)

	bfs := func(x int, y int, ln int) {
		var head, tail int
		que[head] = x*m + y
		head++
		dist[x][y] = 0

		for tail < head {
			u, v := que[tail]/m, que[tail]%m
			tail++

			if color[u][v] == '0' {
				if !ok[dist[u][v]%ln] {
					res[1]++
				}
				ok[dist[u][v]%ln] = true
			}

			for k := 0; k < 4; k++ {
				r, c := u+dd[k], v+dd[k+1]
				x, y := move(r, c)
				if x == u && y == v && dist[r][c] < 0 {
					dist[r][c] = dist[u][v] + 1
					que[head] = r*m + c
					head++
				}
			}
		}

		for i := 0; i < ln; i++ {
			ok[i] = false
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !vis[i][j] {
				// in loop
				u, v := i, j
				var ln int
				for !vis[u][v] {
					ln++
					vis[u][v] = true
					u, v = move(u, v)
				}
				res[0] += ln

				bfs(i, j, ln)
			}
		}
	}

	return res
}
