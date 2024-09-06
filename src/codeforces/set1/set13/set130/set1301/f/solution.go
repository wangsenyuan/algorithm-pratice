package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(reader, m)
	}
	q := readNum(reader)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 4)
	}

	res := solve(k, grid, queries)

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

func solve(k int, grid [][]int, queries [][]int) []int16 {
	n := len(grid)
	m := len(grid[0])

	dist := make([][][]int8, k)

	for i := 0; i < k; i++ {
		dist[i] = make([][]int8, n)
		for j := 0; j < n; j++ {
			dist[i][j] = make([]int8, m)
			for l := 0; l < m; l++ {
				dist[i][j][l] = -1
			}
		}
	}

	pos := make([][]int, k)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			grid[i][j]--
			pos[grid[i][j]] = append(pos[grid[i][j]], i*m+j)
		}
	}

	que := make([]int, n*m)

	ok := make([]bool, k)

	bfs := func(c int) {
		var head, tail int
		for _, p := range pos[c] {
			que[head] = p
			head++
			i, j := p/m, p%m
			dist[c][i][j] = 0
		}

		clear(ok)

		for tail < head {
			u, v := que[tail]/m, que[tail]%m
			tail++

			if !ok[grid[u][v]] {
				ok[grid[u][v]] = true
				for _, p := range pos[grid[u][v]] {
					nu, nv := p/m, p%m
					if dist[c][nu][nv] == -1 {
						dist[c][nu][nv] = dist[c][u][v] + 1
						que[head] = p
						head++
					}
				}
			}

			for i := 0; i < 4; i++ {
				nu, nv := u+dd[i], v+dd[i+1]
				if nu >= 0 && nu < n && nv >= 0 && nv < m && dist[c][nu][nv] == -1 {
					dist[c][nu][nv] = dist[c][u][v] + 1
					que[head] = nu*m + nv
					head++
				}
			}
		}
	}

	for x := 0; x < k; x++ {
		bfs(x)
	}

	ans := make([]int16, len(queries))
	for i, cur := range queries {
		r1, c1, r2, c2 := cur[0]-1, cur[1]-1, cur[2]-1, cur[3]-1
		ans[i] = int16(min(256, abs(r2-r1)+abs(c2-c1)))
		for x := 0; x < k; x++ {
			ans[i] = min(ans[i], 1+int16(dist[x][r1][c1])+int16(dist[x][r2][c2]))
		}
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}
