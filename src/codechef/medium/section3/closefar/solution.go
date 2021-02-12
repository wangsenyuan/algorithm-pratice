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

	n := readNum(reader)

	A := readNNums(reader, n)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		Q[i] = make([]int, 3)
		Q[i][0] = int(s[0])
		pos := readInt(s, 2, &Q[i][1]) + 1
		readInt(s, pos, &Q[i][2])
	}

	res := solve(n, A, E, m, Q)

	var buf bytes.Buffer

	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const N = 35010
const X_I = 100001
const X_O = -1
const INF = 1 << 30

const H = 20
const BLK_SIZE = 350
const X_C = int('C')
const X_F = int('F')

func solve(n int, A []int, E [][]int, m int, Q [][]int) []int64 {
	B := make([]Pair, n)
	for i := 0; i < n; i++ {
		B[i] = Pair{A[i], i}
	}

	sort.Sort(Pairs(B))

	P := make([]int, n)

	for i := 0; i < n; i++ {
		P[B[i].second] = i
	}

	g := NewGraph(n, len(E)*2+10)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	fa := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = make([]int, H)
	}
	D := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)
	seq := make([]int, 2*n+10)
	var time int
	var dfs func(p, u int)
	dfs = func(p, u int) {
		fa[u][0] = p
		for i := 1; i < H; i++ {
			x := fa[u][i-1]
			fa[u][i] = fa[x][i-1]
		}

		in[u] = time
		seq[time] = u
		time++

		for j := g.nodes[u]; j > 0; j = g.next[j] {
			v := g.to[j]
			if v != p {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}

		out[u] = time
		seq[time] = u
		time++
	}

	dfs(0, 0)

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}

		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<uint(i)) >= D[v] {
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

	qq := make([]Query, m)
	for i := 0; i < m; i++ {
		c, u, v := Q[i][0], Q[i][1], Q[i][2]
		u--
		v--
		if in[u] > in[v] {
			u, v = v, u
		}
		p := lca(u, v)
		if u == p || v == p {
			qq[i] = Query{in[u], in[v] + 1, p, c, i}
		} else {
			qq[i] = Query{out[u], in[v] + 1, p, c, i}
		}
	}

	sort.Sort(Queries(qq))

	tree := NewSegTree(n, B)

	cnt := make([]int, n)

	add := func(t int) {
		x := seq[t]
		x = P[x]
		cnt[x]++
		if cnt[x] == 1 {
			tree.Update(x, 1)
		} else {
			tree.Update(x, 0)
		}
	}

	del := func(t int) {
		x := seq[t]
		x = P[x]
		cnt[x]--
		if cnt[x] == 1 {
			tree.Update(x, 1)
		} else {
			tree.Update(x, 0)
		}
	}

	var L, R int

	ans := make([]int64, m)

	for i := 0; i < m; i++ {
		q := qq[i]
		for R < q.r {
			add(R)
			R++
		}
		for R > q.r {
			R--
			del(R)
		}
		for L < q.l {
			del(L)
			L++
		}
		for L > q.l {
			L--
			add(L)
		}

		if q.c == X_F {
			ans[q.i] = tree.GetMaxDiff(P[q.p])
		} else {
			ans[q.i] = tree.GetMinDiff(P[q.p])
		}
	}

	return ans
}

type Node struct {
	l, r    int
	minDiff int64
}

type SegTree struct {
	nodes []Node
	pos   []int
	on    []int
	B     []Pair
}

func NewSegTree(n int, B []Pair) *SegTree {
	nodes := make([]Node, 4*n)
	pos := make([]int, n)
	on := make([]int, n)

	var build func(p, l, r int)

	build = func(p, l, r int) {
		if l == r {
			pos[l] = p
			nodes[p] = Node{X_I, X_O, int64(INF)}
			return
		} else {
			mid := (l + r) >> 1
			build(p<<1, l, mid)
			build((p<<1)|1, mid+1, r)

			x := min(nodes[p<<1].l, nodes[(p<<1)|1].l)
			y := max(nodes[p<<1].r, nodes[(p<<1)|1].r)
			minDiff := min2(nodes[p<<1].minDiff, nodes[(p<<1)|1].minDiff)

			nodes[p] = Node{x, y, minDiff}
		}
	}
	build(1, 0, n-1)

	return &SegTree{nodes, pos, on, B}
}

func (tree *SegTree) Update(x int, s int) {
	nodes := tree.nodes
	on := tree.on
	pos := tree.pos
	B := tree.B

	update := func(x int) {
		if on[x] == s {
			return
		}
		p := pos[x]
		if s == 1 {
			on[x] = 1
			nodes[p].l = x
			nodes[p].r = x
		} else {
			on[x] = 0
			nodes[p].l = X_I
			nodes[p].r = X_O
		}

		p >>= 1

		for p > 0 {
			l := p << 1
			r := l | 1

			nodes[p].l = min(nodes[l].l, nodes[r].l)
			nodes[p].r = max(nodes[l].r, nodes[r].r)
			nodes[p].minDiff = min2(nodes[l].minDiff, nodes[r].minDiff)
			ll := nodes[l].r
			rr := nodes[r].l
			if ll >= 0 && rr < len(B) {
				nodes[p].minDiff = min2(nodes[p].minDiff, int64(B[rr].first-B[ll].first))
			}
			p >>= 1
		}
	}

	update(x)
}

func (tree *SegTree) GetMinDiff(p int) int64 {
	cur := tree.on[p]
	if cur == 0 {
		tree.Update(p, 1)
	}
	res := tree.nodes[1].minDiff

	if cur == 0 {
		tree.Update(p, 0)
	}
	return res
}

func (tree *SegTree) GetMaxDiff(p int) int64 {
	// p 是 lca，需要被重新激活，假设 u ----> p ----> v
	// 那么p会出现2次。所以on[p] 就是 false, 那么p的贡献会被ingore。所以重新激活后，才能得到正确的答案
	cur := tree.on[p]
	if cur == 0 {
		tree.Update(p, 1)
	}
	l := tree.nodes[1].l
	r := tree.nodes[1].r
	if cur == 0 {
		tree.Update(p, 0)
	}

	return int64(tree.B[r].first - tree.B[l].first)
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

func min2(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Query struct {
	l, r, p int
	c       int
	i       int
}

func (this Query) Less(that Query) bool {
	a := this.l / BLK_SIZE
	b := that.l / BLK_SIZE
	if a != b {
		return a < b
	}
	if a%2 == 0 {
		return this.r < that.r
	}
	return this.r > that.r
}

type Queries []Query

func (this Queries) Len() int {
	return len(this)
}

func (this Queries) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Queries) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
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

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
