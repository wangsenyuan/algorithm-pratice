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

	par := make([]int, n)
	d := make([]int, n)
	loc := make([]int, n)

	var dfs func(p, u int)
	dfs = func(p, u int) {
		par[u] = p
		loc[u] = u
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				loc[u] = -1
				d[v] = d[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)

	node := make([]int, n)
	for i := 0; i < n; i++ {
		node[i] = -1
	}

	for i := 0; i < n; i++ {
		if loc[i] != -1 {
			node[loc[i]] = i
		}
	}
	set := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		size[i]++
	}

	var find func(x int) int

	find = func(x int) int {
		if set[x] != x {
			set[x] = find(set[x])
		}
		return set[x]
	}

	union := func(x, y int) {
		px := find(x)
		py := find(y)
		if px != py {
			if size[px] < size[py] {
				px, py = py, px
			}
			set[py] = px
			size[px] += size[py]
		}
	}

	q := make(PriorityQueue, 0, n)

	for i := 0; i < n; i++ {
		if loc[i] > 0 {
			//assassins
			ass := new(Item)
			ass.first = i
			ass.second = i
			ass.priority = d[i]
			heap.Push(&q, ass)
		}
	}
	winAt := -1

	processed := make([]bool, n)

	var process func(q *PriorityQueue, u, p, ancestor int) bool

	process = func(q *PriorityQueue, u, p, ancestor int) bool {
		var contains bool
		var nextAnc = ancestor
		if ancestor == -1 {
			nextAnc = node[u]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || processed[v] {
				continue
			}
			if !contains {
				contains = process(q, v, u, nextAnc)
			}
			processed[v] = true
		}
		if contains && node[u] != -1 {
			loc[node[u]] = -1
			node[u] = -1
		}
		if ancestor != -1 && node[u] != -1 {
			if ancestor != node[u] {
				par[u] = loc[ancestor]
			}
			item := new(Item)
			item.first = node[u]
			item.second = u
			item.priority = d[u]
			heap.Push(q, item)
		}
		return contains || node[u] != -1
	}

	for time := 1; time <= n; time++ {
		nq := make(PriorityQueue, 0, len(q))
		for q.Len() > 0 {
			item := heap.Pop(&q).(*Item)
			loc[item.first] = -1
			node[item.second] = -1
			ntem := new(Item)
			ntem.first = item.first
			ntem.second = par[item.second]
			ntem.priority = d[par[item.second]]
			heap.Push(&nq, ntem)
		}

		prevAss, prevNode := -1, -1

		for nq.Len() > 0 {
			cur := heap.Pop(&nq).(*Item)
			if cur.second == prevNode {
				union(cur.first, prevAss)
			} else {
				prevNode = cur.second
				prevAss = cur.first
				heap.Push(&q, cur)
				loc[cur.first] = cur.second
				node[cur.second] = cur.first
			}
		}

		if node[0] != -1 {
			winAt = time
			break
		}

		for q.Len() > 0 {
			cur := heap.Pop(&q).(*Item)
			if processed[cur.second] {
				continue
			}
			process(&nq, cur.second, par[cur.second], cur.first)
		}

		q = nq
	}

	var ans []int

	for i := 1; i < n; i++ {
		if find(i) == find(node[0]) {
			ans = append(ans, i+1)
		}
	}

	return winAt, ans
}

type Item struct {
	first    int
	second   int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
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
