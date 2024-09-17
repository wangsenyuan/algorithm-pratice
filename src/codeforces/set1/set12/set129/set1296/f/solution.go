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
	questions := make([][]int, m)
	for i := 0; i < m; i++ {
		questions[i] = readNNums(reader, 3)
	}

	f := solve(n, edges, questions)

	if len(f) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	for i := 0; i < n-1; i++ {
		buf.WriteString(fmt.Sprintf("%d ", f[i]))
	}
	fmt.Println(buf.String())
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

const H = 20

const inf = 1 << 30

func solve(n int, edges [][]int, questions [][]int) []int {
	g := NewGraph(n, 2*n)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	dep := make([]int, n)
	fa := make([][]int, n)
	sz := make([]int, n)
	in := make([]int, n)
	var dfs func(p int, u int)

	var timer int

	dfs = func(p int, u int) {
		in[u] = timer
		timer++
		sz[u]++
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
				sz[u] += sz[v]
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
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

	type pair struct {
		first  int
		second int
	}

	occ := make([][]pair, n)

	for _, ques := range questions {
		a, b, g := ques[0]-1, ques[1]-1, ques[2]
		p := lca(a, b)
		if p != a {
			occ[p] = append(occ[p], pair{a, g})
		}
		if p != b {
			occ[p] = append(occ[p], pair{b, g})
		}
	}

	set := NewSegTree(n, 1, max)

	ans := make([]int, n-1)

	var dfs2 func(p int, u int)

	vals := make([]int, n)
	mv := make([][]int, n)
	for i := 0; i < n; i++ {
		mv[i] = make([]int, H)
	}

	dfs2 = func(p int, u int) {
		for i := 1; i < H; i++ {
			mv[u][i] = min(mv[u][i-1], mv[fa[u][i-1]][i-1])
		}
		for _, evt := range occ[u] {
			v, g := evt.first, evt.second
			vals[v] = max(vals[v], g)
			set.Update(in[v], vals[v])
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			j := g.val[i]
			ans[j] = set.Query(in[v], in[v]+sz[v])
			mv[v][0] = ans[j]
			dfs2(u, v)
		}
	}

	for i := 0; i < H; i++ {
		mv[0][i] = inf
	}

	dfs2(0, 0)

	get := func(u int, p int) int {
		res := inf
		if p != u {
			diff := dep[u] - dep[p]
			for i := H - 1; i >= 0; i-- {
				if (diff>>i)&1 == 1 {
					res = min(res, mv[u][i])
					u = fa[u][i]
				}
			}
		}

		return res
	}

	check := func(u int, v int) int {
		p := lca(u, v)
		return min(get(u, p), get(v, p))
	}

	for _, cur := range questions {
		a, b, g := cur[0]-1, cur[1]-1, cur[2]
		if check(a, b) != g {
			return nil
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

type SegTree struct {
	f   func(int, int) int
	arr []int
	iv  int
	n   int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, n*2)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{f, arr, iv, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = v

	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Query(l int, r int) int {
	res := t.iv
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
