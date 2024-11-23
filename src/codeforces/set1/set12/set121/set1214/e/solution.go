package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	d := readNNums(reader, n)
	res := solve(d)
	var buf bytes.Buffer
	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
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

func solve(d []int) [][]int {
	type pair struct {
		first  int
		second int
	}

	n := len(d)
	a := make([]pair, n)
	for i := 0; i < n; i++ {
		a[i] = pair{d[i], i}
	}

	slices.SortFunc(a, func(x, y pair) int {
		return y.first - x.first
	})
	var res [][]int

	x := make([]int, n)
	x[0] = a[0].second*2 + 1
	pos := make([]int, n*2+2)
	pos[x[0]] = 0
	for i := 1; i < n; i++ {
		j := a[i].second
		x[i] = 2*j + 1
		res = append(res, []int{x[i-1], x[i]})
		pos[x[i]] = i
	}

	for _, cur := range a {
		i := cur.second
		di := cur.first
		j := pos[2*i+1]
		k := j - (di - 1)
		if k < 0 {
			k = j + di - 1
		}
		if k < len(x)-1 {
			res = append(res, []int{2*i + 2, x[k]})
		} else {
			res = append(res, []int{2*i + 2, x[k]})
			x = append(x, 2*i+2)
		}
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) prepare() (dist func(int, int) int) {
	n := len(g.nodes)
	dep := make([]int, n)
	h := bits.Len(uint(n))
	fa := make([][]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
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
	dfs(0, 0)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	dist = func(u int, v int) int {
		p := lca(u, v)
		return dep[u] + dep[v] - 2*dep[p]
	}

	return dist
}
