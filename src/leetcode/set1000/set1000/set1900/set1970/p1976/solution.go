package p1976

import (
	"container/heap"
	"math"
)

const MOD = 1000000007

func countPaths(n int, roads [][]int) int {
	g := NewGraph(n, len(roads)*2)
	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	pq := make(PQ, n)
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		tmp := new(Node)
		tmp.index = i
		tmp.pos = i
		tmp.time = math.MaxInt64
		pq[i] = tmp
		nodes[i] = tmp
	}
	nodes[0].count = 1
	pq.update(nodes[0], 0)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Node)
		u := cur.pos

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.cost[i]
			if nodes[v].time > cur.time+int64(w) {
				nodes[v].count = cur.count
				pq.update(nodes[v], cur.time+int64(w))
			} else if nodes[v].time == cur.time+int64(w) {
				nodes[v].count += cur.count
				nodes[v].count %= MOD
			}
		}
	}

	return int(nodes[n-1].count)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cost  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cost = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.cost[g.cur] = w
}

type Node struct {
	index int
	time  int64
	count int64
	pos   int
}

type PQ []*Node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	cur := x.(*Node)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() interface{} {
	old := *(pq)
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}

func (pq *PQ) update(node *Node, time int64) {
	node.time = time
	heap.Fix(pq, node.index)
}
