package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, c0 := readTwoNums(reader)
		c := readNNums(reader, n-1)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, c0, c, edges)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const inf = 1 << 30

func solve(n int, c0 int, c []int, edges [][]int) []int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}

	ans := inf

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		if dist[u] >= ans {
			return
		}
		if vis[u] {
			ans = min(ans, dist[u])
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if dist[u]+1 < dist[v] {
				dist[v] = dist[u] + 1
				dfs(u, v)
			} else {
				// no need to process v
				ans = min(ans, dist[u]+1+dist[v])
			}
		}
	}
	dist[c0-1] = 0
	dfs(-1, c0-1)
	vis[c0-1] = true
	res := make([]int, n-1)

	for i, x := range c {
		x--
		dist[x] = 0
		dfs(-1, x)
		res[i] = ans
		vis[x] = true
	}

	return res
}

func solve1(n int, c0 int, c []int, edges [][]int) []int {
	c0--

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	next := make([]int, n)
	fa := make([]int, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = p
		sz[u] = 1
		next[u] = -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if next[u] < 0 || sz[next[u]] < sz[v] {
					next[u] = v
				}
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, c0)

	head := make([]int, n)
	end := make([]int, n)
	pos := make([]int, n)

	var it int

	var dfs2 func(p int, u int, x int)

	dfs2 = func(p int, u int, x int) {
		head[u] = x
		end[x] = it
		pos[u] = it
		it++
		if next[u] != -1 {
			dfs2(u, next[u], x)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p || next[u] == v {
				continue
			}
			dfs2(u, v, v)
		}
	}

	dfs2(-1, c0, c0)

	tr1 := NewSegTree(n, inf, min)
	tr2 := NewSegTree(n, inf, min)

	process := func(u int) int {
		res := inf
		var dist int

		for u != -1 {
			x := head[u]
			// 找出它前面的
			tmp := tr1.Query(pos[x], pos[u]) + dist + pos[u]
			res = min(res, tmp)
			tmp = tr2.Query(pos[u], end[x]+1) + dist - pos[u]
			res = min(res, tmp)

			tr1.Update(pos[u], dist-pos[u])
			tr2.Update(pos[u], dist+pos[u])
			// go to fa of head
			dist += pos[u] - pos[x] + 1
			u = fa[x]
		}

		return res
	}

	tr1.Update(0, 0)
	tr2.Update(0, 0)

	ans := make([]int, n-1)
	for i, x := range c {
		ans[i] = process(x - 1)
		if i > 0 {
			ans[i] = min(ans[i], ans[i-1])
		}
	}

	return ans
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
	t.arr[p] = t.f(t.arr[p], v)

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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
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
