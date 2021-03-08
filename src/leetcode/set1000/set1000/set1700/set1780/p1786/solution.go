package p1786

import "container/heap"

const INF = 1 << 30
const MOD = 1000000007

func countRestrictedPaths(n int, edges [][]int) int {
	g := NewGraph(n, len(edges)*2+10)

	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	nodes := make([]*Item, n)
	pq := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		nodes[i] = new(Item)
		nodes[i].value = i
		nodes[i].priority = INF
		nodes[i].index = i
		pq[i] = nodes[i]
	}
	heap.Init(&pq)

	pq.update(nodes[n-1], 0)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.value

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.cnt[i]
			if nodes[u].priority+w < nodes[v].priority {
				pq.update(nodes[v], nodes[u].priority+w)
			}
		}
	}

	g2 := NewGraph(n, len(edges)+10)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if nodes[u].priority > nodes[v].priority {
			g2.AddEdge(u, v, 1)
		} else if nodes[u].priority < nodes[v].priority {
			g2.AddEdge(v, u, 1)
		}
	}

	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = -1
	}
	cnt[n-1] = 1

	var dfs func(u int) int

	dfs = func(u int) int {
		if cnt[u] >= 0 {
			return cnt[u]
		}
		cnt[u] = 0
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			cnt[u] += dfs(v)
			if cnt[u] >= MOD {
				cnt[u] -= MOD
			}
		}
		return cnt[u]
	}

	return dfs(0)
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
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
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cnt   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	cnt := make([]int, e)
	return &Graph{nodes, next, to, cnt, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.cnt[g.cur] += w
}
