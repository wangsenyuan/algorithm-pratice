package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, E, Q)

	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

func solve(n int, E [][]int, Q [][]int) []int64 {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	depth := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)

	parent := make([]int, n)

	children := make([]int, n)

	var time int

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		parent[u] = p
		in[u] = time
		time++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				children[u]++
				depth[v] = depth[u] + 1
				dfs(u, v)
			}
		}
		out[u] = time
	}

	dfs(-1, 0)

	dsu := NewDSU(n)

	cur := make([]*Vertex, n)

	pq := make(PQ, 0, n)

	t := NewTree(n)

	addPath := func(x int, y int, v Pair) {
		t.Add(0, 0, n, in[x], v)
		if parent[y] != -1 {
			t.Add(0, 0, n, in[parent[y]], v.Negative())
		}
	}

	getVertexValue := func(v int) Pair {
		return t.Get(0, 0, n, in[v], out[v])
	}

	calculateCost := func(v int, correction int) int {
		res := getVertexValue(v)
		k := (res.first+res.second)/(res.second+1) - 1
		if int64(correction) < k {
			return correction
		}
		return int(k)
	}

	for i := 0; i < n; i++ {
		cnt := children[i]
		addPath(i, i, Pair{int64(cnt), 0})
		if i > 0 {
			cur[i] = NewVertex(calculateCost(i, 200000), depth[i], i)
			heap.Push(&pq, cur[i])
		}
	}

	for i := 0; i < len(Q); i++ {
		q := Q[i]
		heap.Push(&pq, NewVertex(q[1], -1, i))
	}

	erase := func(i int) {
		cur[i].cost = 1 << 30
		heap.Fix(&pq, cur[i].index)
		heap.Pop(&pq)
	}

	ans := make([]int64, len(Q))

	for pq.Len() > 0 {
		z := heap.Pop(&pq).(*Vertex)
		if z.depth == -1 {
			res := getVertexValue(Q[z.id][0] - 1)
			ans[z.id] = res.first - res.second*int64(z.cost)
		} else {
			v := z.id
			p := parent[v]
			top := dsu.FindTop(p)
			val := getVertexValue(v)
			(&val).second++
			(&val).first--
			if top != 0 {
				erase(top)
			}
			addPath(p, top, val)
			if top != 0 {
				cur[top].cost = calculateCost(top, z.cost)
				heap.Push(&pq, cur[top])
			}
			dsu.Merge(v, p)
		}
	}

	return ans
}

type Pair struct {
	first  int64
	second int64
}

func (this Pair) Add(that Pair) Pair {
	return Pair{this.first + that.first, this.second + that.second}
}

func (this Pair) Negative() Pair {
	return Pair{-this.first, -this.second}
}

func (this Pair) Copy() Pair {
	return Pair{this.first, this.second}
}

type Tree struct {
	nodes []Pair
}

func NewTree(n int) *Tree {
	nodes := make([]Pair, 4*n)
	return &Tree{nodes}
}

func (t *Tree) Get(v int, l int, r int, L int, R int) Pair {
	if L >= R {
		return Pair{0, 0}
	}
	if l == L && r == R {
		return t.nodes[v].Copy()
	}
	m := (l + r) / 2
	a := t.Get(2*v+1, l, m, L, min(R, m))
	b := t.Get(2*v+2, m, r, max(L, m), R)
	return a.Add(b)
}

func (t *Tree) Add(v int, l int, r int, pos int, val Pair) {
	t.nodes[v] = t.nodes[v].Add(val)
	if l == r-1 {
		return
	}
	mid := (l + r) / 2
	if pos < mid {
		t.Add(2*v+1, l, mid, pos, val)
	} else {
		t.Add(2*v+2, mid, r, pos, val)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

type Vertex struct {
	cost  int
	depth int
	id    int
	index int
}

func NewVertex(cost int, depth int, id int) *Vertex {
	return &Vertex{cost, depth, id, 0}
}

func (this Vertex) Less(that Vertex) bool {
	if this.cost != that.cost {
		return this.cost > that.cost
	}
	if this.depth != that.depth {
		return this.depth > that.depth
	}
	return this.id < that.id
}

type PQ []*Vertex

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Less(*pq[j])
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	v := x.(*Vertex)
	v.index = len(*pq)
	*pq = append(*pq, v)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

type DSU struct {
	arr []int
	top []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	top := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		top[i] = i
	}
	return &DSU{arr, top}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) FindTop(x int) int {
	return this.top[this.Find(x)]
}

func (this *DSU) Merge(x, p int) {
	this.arr[x] = p
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	e++
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
