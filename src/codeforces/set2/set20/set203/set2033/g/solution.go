package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	qs := make([][]int, m)
	for i := range m {
		qs[i] = readNNums(reader, 2)
	}
	return solve(n, edges, qs)
}

func solve(n int, edges [][]int, qs [][]int) []int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	type pair struct {
		first  int
		second int
	}

	dep := make([]int, n)
	far := make([][]pair, n)

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		ans := dep[u]
		far[u] = []pair{{dep[u], u}, {-inf, -1}}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				tmp := dfs(u, v)
				if tmp >= far[u][0].first {
					far[u][1] = far[u][0]
					far[u][0] = pair{tmp, v}
				} else if tmp >= far[u][1].first {
					far[u][1] = pair{tmp, v}
				}
				ans = max(ans, tmp)
			}
		}
		return ans
	}

	dfs(-1, 0)

	at := make([][]int, n)
	for i, cur := range qs {
		v := cur[0] - 1
		at[v] = append(at[v], i)
	}

	ans := make([]int, len(qs))

	tr := NewSegTree(n)

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		for _, i := range at[u] {
			k := qs[i][1]
			// 子树u的最远距离，不需要额外的花费
			ans[i] = far[u][0].first - dep[u]
			l := max(0, dep[u]-k)
			tmp := tr.Get(l, dep[u])
			ans[i] = max(ans[i], tmp+dep[u])
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				w := far[u][0]
				if w.second == v {
					w = far[u][1]
				}
				if w.second < 0 {
					// 没有第二个子树
					tr.Update(dep[u], -dep[u])
				} else {
					tr.Update(dep[u], w.first-2*dep[u])
				}
				dfs2(u, v)
			}
		}
	}

	tr.Update(0, 0)

	dfs2(-1, 0)

	return ans
}

const inf = 1 << 30

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := range len(arr) {
		arr[i] = -inf
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	res := -inf
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = max(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
