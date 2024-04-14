package p3112

import "container/heap"

const oo = 1 << 60

func minimumTime(n int, edges [][]int, disappear []int) []int {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
		it := new(Item)
		it.id = i
		it.priority = oo
		it.index = i
		items[i] = it
		pq[i] = it
	}

	heap.Init(&pq)

	pq.update(items[0], 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority >= disappear[it.id] {
			continue
		}
		if ans[it.id] < 0 {
			ans[it.id] = it.priority
		}
		u := it.id
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if it.priority+w < items[v].priority {
				pq.update(items[v], it.priority+w)
			}
		}
	}

	return ans
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
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
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
	priority int
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
	return item
}

func (pq *PriorityQueue) update(item *Item, value int) {
	item.priority = value
	heap.Fix(pq, item.index)
}
