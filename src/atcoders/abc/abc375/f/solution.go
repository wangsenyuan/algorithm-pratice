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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int64 {
	n, m, q := readThreeNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			queries[i] = make([]int, 2)
			queries[i][0] = 1
		} else {
			queries[i] = make([]int, 3)
			queries[i][0] = 2
		}
		pos := 2
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}

	return solve(n, edges, queries)
}

const inf = 1 << 60

func solve(n int, edges [][]int, queries [][]int) []int64 {
	m := len(edges)
	closedAt := make([]int, m)
	for i := 0; i < m; i++ {
		closedAt[i] = -1
	}
	for i, cur := range queries {
		if cur[0] == 1 {
			j := cur[1] - 1
			closedAt[j] = i
		}
	}

	var e1 [][]int
	for i := 0; i < m; i++ {
		if closedAt[i] < 0 {
			e1 = append(e1, edges[i])
		}
	}
	dist := calcDist(n, e1)

	addEdge := func(i int) {
		e := edges[i]
		u, v, w := e[0]-1, e[1]-1, int64(e[2])

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][u]+dist[v][j]+w)
				dist[i][j] = min(dist[i][j], dist[i][v]+dist[u][j]+w)
			}
		}
	}

	k := len(queries)
	var ans []int64

	for i := k - 1; i >= 0; i-- {
		cur := queries[i]
		if cur[0] == 1 {
			addEdge(cur[1] - 1)
		} else {
			u, v := cur[1]-1, cur[2]-1
			if dist[u][v] == inf {
				ans = append(ans, -1)
			} else {
				ans = append(ans, dist[u][v])
			}
		}
	}

	reverse(ans)

	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func reverse(arr []int64) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func calcDist(n int, edges [][]int) [][]int64 {

	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		it := new(Item)
		it.id = i
		it.priority = inf
		items[i] = it
	}

	g := NewGraph(n, len(edges)*2)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	f := func(s int, dist []int64) {
		pq := make(PriorityQueue, n)
		for i := 0; i < n; i++ {
			items[i].priority = inf
			items[i].index = i
			pq[i] = items[i]
		}
		pq.update(items[s], 0)
		for pq.Len() > 0 {
			it := heap.Pop(&pq).(*Item)
			u := it.id
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := int64(g.val[i])
				if items[v].index >= 0 && items[v].priority > it.priority+w {
					pq.update(items[v], it.priority+w)
				}
			}
		}

		for i := 0; i < n; i++ {
			dist[i] = items[i].priority
		}
	}

	dist := make([][]int64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int64, n)
		f(i, dist[i])
	}

	return dist
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	val := make([]int, e+10)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int64
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, v int64) {
	it.priority = v
	heap.Fix(pq, it.index)
}
