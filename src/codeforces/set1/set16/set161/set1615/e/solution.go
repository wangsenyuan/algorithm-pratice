package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
)

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
	var buf bytes.Buffer
	tc := 1

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, m)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(n int, E [][]int, k int) int64 {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	H := make([]int, n)
	P := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)

	var dfs func(p int, u int)

	var time int

	ord := make([]int, n)
	dfs = func(p int, u int) {
		P[u] = p
		in[u] = time
		ord[time] = u
		time++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				H[v] = H[u] + 1
				dfs(u, v)
			}
		}
		out[u] = time
	}

	dfs(-1, 0)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = H[ord[i]]
	}

	root := BuildTree(arr)

	var best int64 = math.MinInt64

	var w int

	pq := make(PQ, 0, n)

	mark := make([]bool, n)

	heap.Push(&pq, root.Get(0, n-1))

	for r := 1; r <= k; r++ {
		if pq.Len() > 0 {
			cur := heap.Pop(&pq).(Pair)
			// cur.first is the height, cur.second is the pos
			u := ord[cur.second]

			for P[u] >= 0 && !mark[P[u]] {
				p := P[u]
				mark[p] = true

				w++

				for i := g.nodes[p]; i > 0; i = g.next[i] {
					v := g.to[i]
					if v == P[p] || v == u {
						continue
					}
					h := root.Get(in[v], in[v]).first
					root.Update(in[v], out[v]-1, -h)
					tmp := root.Get(in[v], out[v]-1)
					heap.Push(&pq, tmp)
				}

				u = p
			}
		}
		// r * (n - r) - b * (n - b)

		b := max(min(n-w-r, n/2), 0)

		tmp := int64(r)*int64(n-r) - int64(b)*int64(n-b)

		best = max2(best, tmp)
	}

	return best
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
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

const INF = 1e9

type Pair struct {
	first  int
	second int
}

func max_pair(a, b Pair) Pair {
	if a.first > b.first {
		return a
	}
	return b
}

func (this Pair) addFirst(v int) Pair {
	return Pair{this.first + v, this.second}
}

type Tree struct {
	key  []Pair
	lazy []int
	sz   int
}

func BuildTree(arr []int) *Tree {
	n := len(arr)
	key := make([]Pair, 4*n)
	lazy := make([]int, 4*n)
	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		if l == r {
			key[i] = Pair{arr[l], l}
			return
		}
		mid := (l + r) / 2
		build(i*2, l, mid)
		build(i*2+1, mid+1, r)
		key[i] = max_pair(key[2*i], key[2*i+1])
	}

	build(1, 0, n-1)

	return &Tree{key, lazy, n}
}

func (t *Tree) updateLazy(i int, v int) {
	t.lazy[i] += v
	t.key[i] = t.key[i].addFirst(v)
}

func (t *Tree) push(i int, l int, r int) {
	if l < r && t.lazy[i] != 0 {
		t.updateLazy(2*i, t.lazy[i])
		t.updateLazy(2*i+1, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		t.push(i, l, r)

		if L <= l && r <= R {
			t.updateLazy(i, v)
			return
		}
		mid := (l + r) / 2
		loop(2*i, l, mid)
		loop(2*i+1, mid+1, r)
		t.key[i] = max_pair(t.key[i*2], t.key[2*i+1])
	}

	loop(1, 0, t.sz-1)
}

func (t *Tree) Get(L int, R int) Pair {
	var loop func(i int, l int, r int) Pair
	loop = func(i int, l int, r int) Pair {
		if R < l || r < L {
			return Pair{-INF, -1}
		}
		t.push(i, l, r)
		if L <= l && r <= R {
			return t.key[i]
		}
		mid := (l + r) / 2
		a := loop(i*2, l, mid)
		b := loop(i*2+1, mid+1, r)
		return max_pair(a, b)
	}

	return loop(1, 0, t.sz-1)
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

type PQ []Pair

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].first > pq[j].first
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	it := x.(Pair)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}
