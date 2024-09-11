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
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 5)
	}

	res := solve(n, edges, queries)

	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const H = 20

func solve(n int, edges [][]int, queries [][]int) []bool {
	g := NewGraph(n+1, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dep := make([]int, n+1)
	fa := make([][]int, n+1)
	fa[0] = make([]int, H)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = make([]int, H)
		fa[u][0] = p
		for i := 1; i < H; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 1)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	getDist := func(u int, v int) int {
		p := lca(u, v)
		return dep[u] + dep[v] - 2*dep[p]
	}

	check := func(x, y, a, b, k int) bool {
		dab := getDist(a, b)
		if k >= dab && (k-dab)%2 == 0 {
			return true
		}
		dax := getDist(a, x)
		dby := getDist(b, y)
		dab = dax + dby + 1
		if k >= dab && (k-dab)%2 == 0 {
			return true
		}

		day := getDist(a, y)
		dbx := getDist(b, x)
		dab = day + dbx + 1
		if k >= dab && (k-dab)%2 == 0 {
			return true
		}

		return false
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		ans[i] = check(cur[0], cur[1], cur[2], cur[3], cur[4])
	}

	return ans
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
