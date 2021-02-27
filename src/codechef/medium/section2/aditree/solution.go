package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	solver := NewSolver(n, E)

	q := readNum(reader)

	var buf bytes.Buffer

	for q > 0 {
		q--
		a, b := readTwoNums(reader)
		res := solver.TurnOn(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

const H = 20

type Solver struct {
	n          int
	P          [][]int
	D          []int
	chainHead  []int
	chainInd   []int
	treePos    []int
	firstChild []int
	tree       *SegTree
}

func NewSolver(n int, E [][]int) *Solver {
	g := NewGraph(n, len(E)*2+10)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}

	P := make([][]int, H)
	D := make([]int, n)
	SZ := make([]int, n)
	var dfs func(p, u int)

	dfs = func(p, u int) {
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			if P[u][i-1] >= 0 {
				P[u][i] = P[P[u][i-1]][i-1]
			} else {
				P[u][i] = -1
			}
		}
		SZ[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
				SZ[u] += SZ[v]
			}
		}
	}

	dfs(-1, 0)

	var hld func(p, u int)
	chainHead := make([]int, n)
	chainInd := make([]int, n)
	treePos := make([]int, n)
	hc := make([]int, n)
	for i := 0; i < n; i++ {
		chainHead[i] = -1
		hc[i] = -1
	}
	var chainNo int
	var pos int
	hld = func(p, u int) {
		if chainHead[chainNo] < 0 {
			chainHead[chainNo] = u
		}
		chainInd[u] = chainNo
		treePos[u] = pos
		pos++
		ch := -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if ch < 0 || SZ[v] > SZ[ch] {
					ch = v
				}
			}
		}
		if ch < 0 {
			return
		}
		hc[u] = ch
		hld(u, ch)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p || v == ch {
				continue
			}
			chainNo++
			hld(u, v)
		}
	}

	hld(-1, 0)

	tree := NewSegTree(n)

	return &Solver{n, P, D, chainHead, chainInd, treePos, hc, tree}
}

func (solver *Solver) lca(u, v int) int {

	if solver.D[u] < solver.D[v] {
		u, v = v, u
	}

	for i := H - 1; i >= 0; i-- {
		if solver.D[u]-(1<<uint(i)) >= solver.D[v] {
			u = solver.P[u][i]
		}
	}

	if u == v {
		return u
	}

	for i := H - 1; i >= 0; i-- {
		if solver.P[u][i] != solver.P[v][i] {
			u = solver.P[u][i]
			v = solver.P[v][i]
		}
	}
	return solver.P[u][0]
}

func (solver *Solver) updateUp(a int, p int) {

	for solver.chainInd[a] != solver.chainInd[p] {
		h := solver.chainHead[solver.chainInd[a]]
		solver.tree.flip(solver.treePos[h], solver.treePos[a])
		a = solver.P[h][0]
	}

	if a != p {
		solver.tree.flip(solver.treePos[solver.firstChild[p]], solver.treePos[a])
	}
}

func (solver *Solver) TurnOn(a, b int) int {
	a--
	b--
	p := solver.lca(a, b)
	solver.updateUp(a, p)
	solver.updateUp(b, p)
	return solver.tree.query(0, solver.tree.n-1)
}

type SegTree struct {
	lazy []int
	sum  []int
	n    int
}

func NewSegTree(n int) *SegTree {
	m := 1
	for m < n {
		m *= 2
	}
	lazy := make([]int, m<<1)
	sum := make([]int, m<<1)
	return &SegTree{lazy, sum, m}
}

func (tree *SegTree) flip(l, r int) {
	lazy := tree.lazy
	sum := tree.sum

	var loop func(i, ll, rr int)

	loop = func(i, ll, rr int) {
		if lazy[i] > 0 {
			sum[i] = rr - ll + 1 - sum[i]
			if ll != rr {
				lazy[i<<1] ^= 1
				lazy[(i<<1)|1] ^= 1
			}
			lazy[i] ^= 1
		}
		if r < ll || rr < l || ll > rr {
			return
		}

		if l <= ll && rr <= r {
			sum[i] = rr - ll + 1 - sum[i]
			if ll != rr {
				lazy[i<<1] ^= 1
				lazy[(i<<1)|1] ^= 1
			}
			return
		}

		mid := (ll + rr) >> 1
		loop(i<<1, ll, mid)
		loop((i<<1)|1, mid+1, rr)
		sum[i] = sum[i<<1] + sum[(i<<1)|1]
	}
	loop(1, 0, tree.n-1)
}

func (tree *SegTree) query(l, r int) int {
	lazy := tree.lazy
	sum := tree.sum

	var loop func(i, ll, rr int) int

	loop = func(i, ll, rr int) int {
		if lazy[i] > 0 {
			sum[i] = rr - ll + 1 - sum[i]
			if ll != rr {
				lazy[i<<1] ^= 1
				lazy[(i<<1)|1] ^= 1
			}
			lazy[i] ^= 1
		}
		if ll > rr || r < ll || rr < l {
			return 0
		}
		if l <= ll && rr <= r {
			return sum[i]
		}
		mid := (ll + rr) >> 1
		return loop(i<<1, ll, mid) + loop((i<<1)|1, mid+1, rr)
	}

	return loop(1, 0, tree.n-1)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, sz int) *Graph {
	nodes := make([]int, n)
	next := make([]int, sz)
	to := make([]int, sz)
	return &Graph{nodes, next, to, 1}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
