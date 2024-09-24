package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)[:m]
	}

	res := solve(grid)

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

const inf = 1 << 60

func solve(grid []string) int {
	n := len(grid)
	m := len(grid[0])

	que := make([]int, n*m)

	bfs := func(src [][]int) [][]int {
		dist := make([][]int, n)
		for i := 0; i < n; i++ {
			dist[i] = make([]int, m)
			for j := 0; j < m; j++ {
				dist[i][j] = -1
			}
		}
		var head, tail int
		for _, cell := range src {
			x, y := cell[0], cell[1]
			dist[x][y] = 0
			que[head] = x*m + y
			head++
		}
		for tail < head {
			u, v := que[tail]/m, que[tail]%m
			tail++
			for i := 0; i < 4; i++ {
				r, c := u+dd[i], v+dd[i+1]
				if r >= 0 && r < n && c >= 0 && c < m && dist[r][c] < 0 && grid[r][c] != '#' {
					dist[r][c] = dist[u][v] + 1
					que[head] = r*m + c
					head++
				}
			}
		}
		return dist
	}

	states := make([][][]int, 3)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '.' || grid[i][j] == '#' {
				continue
			}
			v := int(grid[i][j] - '0')
			v--
			states[v] = append(states[v], []int{i, j})
		}
	}

	best := inf

	dists := make([][][]int, 3)
	for i := 0; i < 3; i++ {
		dists[i] = bfs(states[i])

		da := []int{inf, inf, inf}
		// 如果i在中间
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			for _, cur := range states[j] {
				x, y := cur[0], cur[1]
				if dists[i][x][y] < 0 {
					// 到达不了这里
					return -1
				}
				da[j] = min(da[j], dists[i][x][y])
			}
		}

		da[i] = 0
		best = min(best, da[0]+da[1]+da[2]-2)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '.' {
				// reachable
				if dists[0][i][j] >= 0 && dists[1][i][j] >= 0 && dists[2][i][j] >= 0 {
					best = min(best, dists[0][i][j]+dists[1][i][j]+dists[2][i][j]-2)
				}
			}
		}
	}

	return best
}
