package p3419

import "container/heap"

const inf = 1 << 60

func minMaxWeight(n int, edges [][]int, threshold int) int {

	// 主要还有一个节点至多有threshold条出边
	// 如果是一棵树，那么除了节点0，所有点只有一条边，需要到达它这里
	// 那么似乎不需要关心threshold的限制
	// 且删除的边越多越好。
	items := make([]*Item, n)
	for i := range n {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = -2
		items[i] = it
	}

	g := NewGraph(n, len(edges))

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		// reverse it
		g.AddEdge(v, u, w)
	}

	var ans int

	pq := make(PriorityQueue, 0, n)

	items[0].priority = 0
	heap.Push(&pq, items[0])

	marked := make([]bool, n)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		ans = max(ans, it.priority)
		u := it.id
		marked[u] = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if !marked[v] && w < items[v].priority {
				if items[v].priority == inf {
					items[v].priority = w
					heap.Push(&pq, items[v])
				} else {
					pq.update(items[v], w)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !marked[i] {
			return -1
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
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
