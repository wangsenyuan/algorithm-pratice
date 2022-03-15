package p2203

import (
	"container/heap"
	"math"
)

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	g := NewGraph(n, len(edges)+1)
	rg := NewGraph(n, len(edges)+1)

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		g.AddEdge(u, v, w)
		rg.AddEdge(v, u, w)
	}
	dist1 := bfs(g, n, src1)
	dist2 := bfs(g, n, src2)
	dist3 := bfs(rg, n, dest)
	var res int64 = INF
	for i := 0; i < n; i++ {
		a := dist1[i]
		b := dist2[i]
		c := dist3[i]
		if a == INF || b == INF || c == INF {
			continue
		}
		if a+b+c < res {
			res = a + b + c
		}
	}
	if res == INF {
		return -1
	}
	return res
}

const INF = math.MaxInt64 >> 1

func bfs(g *Graph, n int, s int) []int64 {
	pq := make(PriorityQueue, n)
	items := make([]*Item, n)

	for i := 0; i < n; i++ {
		item := new(Item)
		item.value = i
		item.priority = INF
		item.index = i
		items[i] = item
		pq[i] = item
	}

	items[s].priority = 0
	heap.Init(&pq)

	dist := make([]int64, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[s] = 0

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.value
		dist[u] = cur.priority
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := int64(g.weight[i])
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
				items[v].priority = dist[v]
				pq.update(items[v])
			}
		}
	}

	return dist
}

type Item struct {
	value    int
	priority int64
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
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
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}

func (pq *PriorityQueue) update(item *Item) {
	heap.Fix(pq, item.index)
}

type Graph struct {
	node   []int
	next   []int
	to     []int
	weight []int
	cur    int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.weight = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.weight[g.cur] = w
}
