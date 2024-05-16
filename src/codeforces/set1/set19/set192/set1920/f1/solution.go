package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, q := readThreeNums(reader)
	mat := make([]string, n)

	for i := 0; i < n; i++ {
		mat[i] = readString(reader)[:m]
	}

	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(n, m, mat, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(n int, m int, mat []string, queries [][]int) []int {
	// 先计算每个格子的最短safe距离
	// 然后对于每个query进行二分查找，看是否能在safe距离条件下，形成一个路径
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	que := make([]int, n*m)
	var front, tail int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 'v' {
				que[front] = i*m + j
				front++
				dist[i][j] = 0
			}
		}
	}

	for tail < front {
		pos := front
		for tail < pos {
			u, v := que[tail]/m, que[tail]%m
			tail++
			for i := 0; i < 4; i++ {
				x, y := u+dd[i], v+dd[i+1]
				if x >= 0 && x < n && y >= 0 && y < m && dist[x][y] < 0 {
					dist[x][y] = dist[u][v] + 1
					que[front] = x*m + y
					front++
				}
			}
		}
	}

	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}

	var dfs func(id int) bool

	dfs = func(id int) bool {
		u, v := id/m, id%m
		vis[u][v] = true
		if mat[u][v] == '#' {
			return true
		}

		// 斜角线没有处理
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				x, y := u+dx, v+dy
				if x >= 0 && x < n && y >= 0 && y < m && !vis[x][y] {
					if dfs(x*m + y) {
						return true
					}
				}
			}
		}

		return false
	}

	check := func(pos []int, d int) bool {
		// 能否在d的距离上安全航行
		if dist[pos[0]][pos[1]] < d {
			return false
		}
		front, tail = 0, 0

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				vis[i][j] = false
			}
		}

		vis[pos[0]][pos[1]] = true
		que[front] = pos[0]*m + pos[1]
		front++

		for tail < front {
			u, v := que[tail]/m, que[tail]%m
			tail++
			for i := 0; i < 4; i++ {
				x, y := u+dd[i], v+dd[i+1]
				if x >= 0 && x < n && y >= 0 && y < m && !vis[x][y] && dist[x][y] >= d && mat[x][y] != '#' {
					vis[x][y] = true
					que[front] = x*m + y
					front++
				}
			}
		}

		// 然后检查从最外面是否能到达目的地
		for i := 0; i < n; i++ {
			if !vis[i][0] {
				if dfs(i * m) {
					return false
				}
			}
			if !vis[i][m-1] {
				if dfs(i*m + m - 1) {
					return false
				}
			}
		}

		for j := 0; j < m; j++ {
			if !vis[0][j] {
				if dfs(j) {
					return false
				}
			}
			if !vis[n-1][j] {
				if dfs((n-1)*m + j) {
					return false
				}
			}
		}

		return true
	}

	find := func(pos []int) int {
		l, r := 0, dist[pos[0]][pos[1]]
		if check(pos, r) {
			return r
		}
		for l < r {
			mid := (l + r) / 2
			if check(pos, mid) {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return r - 1
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		cur[0]--
		cur[1]--
		ans[i] = find(cur)
	}

	return ans
}
