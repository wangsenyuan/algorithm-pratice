package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		at, res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d %d\n", len(res), at))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, E [][]int) (int, []int) {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}

	ti := make([][]int, n)
	dep := make([]int, n)

	var dfs1 func(p int, u int, d int)

	dfs1 = func(p int, u int, d int) {
		dep[u] = d
		leaf := true
		a := 1 << 30
		b := 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			leaf = false
			dfs1(u, v, d+1)
			a = min(a, ti[v][1]+1)
			b = max(b, ti[v][1]+1)
		}
		if leaf {
			ti[u] = []int{0, 0}
		} else {
			if a == b {
				ti[u] = []int{a, a}
			} else {
				ti[u] = []int{a, a + 1}
			}
		}
	}

	dfs1(-1, 0, 0)

	X := make(PriorityQueue, 0, n)

	var dfs2 func(p int, u int)

	var flag bool

	gf := NewGraph(n, 2*n)

	dfs2 = func(p int, u int) {
		if X.Len() > 0 && X[0].priority == ti[u][1] {
			gf.AddEdge(u, X[0].pos)
			gf.AddEdge(X[0].pos, u)
		} else if flag {
			gf.AddEdge(u, p)
			gf.AddEdge(p, u)
		}
		item := &Item{depth: dep[u], pos: u, priority: ti[u][0]}
		heap.Push(&X, item)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				flag = ti[u][0] == ti[u][1]
				dfs2(u, v)
			}
		}
		X.update(item, -1)
		heap.Pop(&X)
	}

	dfs2(-1, 0)

	mi := 1 << 30
	for i := g.nodes[0]; i > 0; i = g.next[i] {
		u := g.to[i]
		if ti[u][1] < mi {
			mi = ti[u][1]
		}
	}

	var ans []int
	vis := make([]bool, n)

	var dfs3 func(u int, x int)

	dfs3 = func(u int, x int) {
		if vis[u] {
			return
		}
		vis[u] = true
		if gf.size[u] == 1 && u != x || gf.size[u] == 0 || gf.size[u] == 1 && gf.to[gf.nodes[u]] == 0 {
			ans = append(ans, u+1)
		}

		for i := gf.nodes[u]; i > 0; i = gf.next[i] {
			v := gf.to[i]
			if v != 0 {
				dfs3(v, x)
			}
		}
	}

	for i := g.nodes[0]; i > 0; i = g.next[i] {
		u := g.to[i]
		if mi == ti[u][1] {
			dfs3(u, u)
		}
	}

	sort.Ints(ans)

	return mi + 1, ans
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

type Item struct {
	depth    int
	pos      int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority || pq[i].priority == pq[j].priority && pq[i].depth < pq[j].depth
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	size  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	g.size = make([]int, n)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.size[u]++
}
