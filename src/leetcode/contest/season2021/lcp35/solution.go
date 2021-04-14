package lcp35

import "container/heap"

const INF = 1 << 30

func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	g := NewGraph(n, len(paths)*2)

	for _, path := range paths {
		u, v, w := path[0], path[1], path[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}
	pq := make(PriorityQueue, 0, (cnt+1)*n)
	dp := make([][]*Item, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]*Item, cnt+1)
		for j := 0; j <= cnt; j++ {
			dp[i][j] = NewItem(j*n+i, INF)
			heap.Push(&pq, dp[i][j])
		}
	}
	for j := 0; j <= cnt; j++ {
		pq.update(dp[start][j], j*charge[start])
	}

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		if cur.cost >= INF {
			break
		}
		// u is the node, j is the power
		u, j := cur.value%n, cur.value/n

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if j < w {
				// can't reach v with power j
				continue
			}
			// jj is the power left
			jj := j - w
			for k := 0; k+jj <= cnt; k++ {
				if dp[v][k+jj].cost > dp[u][j].cost+w+k*charge[v] {
					pq.update(dp[v][k+jj], dp[u][j].cost+w+k*charge[v])
				}
			}
		}
	}
	best := INF
	for j := 0; j <= cnt; j++ {
		best = min(best, dp[end][j].cost)
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	value int // The value of the item; arbitrary.
	cost  int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem(value, cost int) *Item {
	item := new(Item)
	item.value = value
	item.cost = cost
	return item
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cost < pq[j].cost
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
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, cost int) {
	item.cost = cost
	heap.Fix(pq, item.index)
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
