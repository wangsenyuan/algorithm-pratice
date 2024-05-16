package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, -1, 0, 1}

const oo = 1 << 30

const H = 20

func solve(n int, m int, mat []string, queries [][]int) []int {
	dist := getDistance(n, m, mat)

	encode := func(i int, j int, b int) int {
		return i*m*2 + j*2 + b
	}

	vx, vy := getVolcano(mat)

	var edges [][]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 2; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if ni >= 0 && ni < n && nj >= 0 && nj < m && mat[i][j] != '#' && mat[ni][nj] != '#' {
					w := min(dist[i][j], dist[ni][nj])
					if i == vx && ni == vx-1 && j > vy {
						edges = append(edges, []int{encode(i, j, 0), encode(ni, nj, 1), w})
						edges = append(edges, []int{encode(i, j, 1), encode(ni, nj, 0), w})
					} else {
						edges = append(edges, []int{encode(i, j, 0), encode(ni, nj, 0), w})
						edges = append(edges, []int{encode(i, j, 1), encode(ni, nj, 1), w})
					}
				}
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] > edges[j][2]
	})
	// a tree
	g := NewGraph(n*m*4, 4*n*m)
	id := n * m * 2

	set := make([]int, n*m*4)

	val := make([]int, n*m*4)

	for i := 0; i < len(set); i++ {
		set[i] = i
	}
	var find func(u int) int

	find = func(u int) int {
		if set[u] != u {
			set[u] = find(set[u])
		}
		return set[u]
	}

	merge := func(u int, v int, w int) {
		u = find(u)
		v = find(v)
		if u == v {
			return
		}
		set[u] = id
		set[v] = id
		val[id] = w
		g.AddEdge(id, u)
		g.AddEdge(id, v)
		id++
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		merge(u, v, w)
	}

	up := make([][]int, id)
	dep := make([]int, id)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		up[u] = make([]int, H)
		up[u][0] = p
		for i := 1; i < H; i++ {
			up[u][i] = up[up[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dep[v] = dep[u] + 1
			dfs(u, v)
		}
	}

	for i := id - 1; i >= 0; i-- {
		if len(up[i]) == 0 {
			dfs(i, i)
		}
	}

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
				u = up[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if up[u][i] != up[v][i] {
				u = up[u][i]
				v = up[v][i]
			}
		}
		return up[u][0]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		x, y := cur[0]-1, cur[1]-1
		p := lca(encode(x, y, 0), encode(x, y, 1))
		ans[i] = val[p]
	}

	return ans
}

func getVolcano(mat []string) (x, y int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == '#' {
				return i, j
			}
		}
	}
	return
}

func getDistance(n int, m int, mat []string) [][]int {
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
				x, y := u+dx[i], v+dy[i]
				if x >= 0 && x < n && y >= 0 && y < m && dist[x][y] < 0 {
					dist[x][y] = dist[u][v] + 1
					que[front] = x*m + y
					front++
				}
			}
		}
	}
	return dist
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
