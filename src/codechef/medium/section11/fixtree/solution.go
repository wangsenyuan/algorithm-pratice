package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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
	P := readNNums(reader, n-1)
	V := readNNums(reader, n)
	q := readNum(reader)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(n, V, P, Q)
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
}

const H = 20

func solve(n int, V []int, parent []int, Q [][]int) []int64 {
	// 麻烦的是  2 X, K => V[x] = K
	L := make([]int, n)
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, H)
		if i > 0 {
			parent[i-1]--
		}
	}

	updateParent := func(p, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			if P[u][i-1] >= 0 {
				P[u][i] = P[P[u][i-1]][i-1]
			}
		}
	}

	marked := make([]bool, n)
	in := make([]int, n)
	out := make([]int, n)
	g := NewGraph(n, n)
	tree := NewFenwick(n)

	var dfs func(p, u int, t *int)

	build := func() {
		g.Reset()

		for i := 0; i < n; i++ {
			marked[i] = false
			if i > 0 {
				P[i][0] = parent[i-1]
				for j := 1; j < H; j++ {
					if P[i][j-1] >= 0 {
						P[i][j] = P[P[i][j-1]][j-1]
					} else {
						P[i][j] = -1
					}
				}
				g.AddEdge(parent[i-1], i)
			}
		}

		dfs(-1, 0, new(int))

		tree.Reset()

		for i := 0; i < n; i++ {
			tree.Update(in[i], int64(V[i]))
			tree.Update(out[i], -int64(V[i]))
		}
	}

	dfs = func(p, u int, t *int) {
		in[u] = *t
		*t++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				L[v] = L[u] + 1
				dfs(u, v, t)
			}
		}
		out[u] = *t
	}

	lca := func(u, v int) int {
		if L[u] < L[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if (L[u] - (1 << uint(i))) >= L[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	blk := 3 * int(math.Sqrt(float64(len(Q))))
	// blk := 1
	var res []int64

	for i, cur := range Q {
		if i%blk == 0 {
			build()
		}

		if cur[0] == 1 {
			x, y := cur[1], cur[2]
			x--
			y--
			parent[x-1] = y
			updateParent(y, x)
			L[x] = L[y] + 1
			marked[x] = true
		} else if cur[0] == 2 {
			x, y := cur[1], cur[2]
			x--
			tree.Update(in[x], int64(y-V[x]))
			tree.Update(out[x], int64(V[x]-y))
			V[x] = y
		} else {
			x, y := cur[1], cur[2]
			x--
			y--
			l := lca(x, y)
			var ans int64
			for x != l && marked[x] {
				ans += int64(V[x])
				x = P[x][0]
			}
			for y != l && marked[y] {
				ans += int64(V[y])
				y = P[y][0]
			}
			ans += tree.GetSum(in[x]) + tree.GetSum(in[y]) - 2*tree.GetSum(in[l])
			ans += int64(V[l])
			res = append(res, ans)
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
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) Reset() {
	for i := 0; i < len(g.nodes); i++ {
		g.nodes[i] = 0
	}
	g.cur = 0
}

type Fenwick struct {
	arr []int64
	sz  int
}

func NewFenwick(n int) *Fenwick {
	tree := new(Fenwick)
	tree.arr = make([]int64, n+1)
	tree.sz = n
	return tree
}

func (t *Fenwick) Update(p int, v int64) {
	arr := t.arr
	p++
	for p <= t.sz {
		arr[p] += v
		p += p & -p
	}
}

func (t *Fenwick) GetSum(p int) int64 {
	p++
	var res int64

	for p > 0 {
		res += t.arr[p]
		p -= p & -p
	}
	return res
}

func (t *Fenwick) GetRangeSum(l, r int) int64 {
	return t.GetSum(r) - t.GetSum(l-1)
}

func (t *Fenwick) Reset() {
	for i := 0; i <= t.sz; i++ {
		t.arr[i] = 0
	}
}
