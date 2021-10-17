package p2044

import (
	"container/heap"
	"math"
)

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// n <= 10000
	// e <= 20000
	g := NewGraph(n, len(edges)*2+2)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	next := func(cur int) int {
		x := cur / change
		if x%2 == 0 {
			// it is green now
			return cur + time
		}
		// it is red now
		return (x+1)*change + time
	}

	dist[0][0] = 0

	pq := make(PriorityQueue, 0, n*2)

	heap.Push(&pq, NewItem(0, 0))

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u, t := cur.value/2, cur.value%2
		if dist[u][t] != cur.priority {
			continue
		}
		// t is the time reaching at node u

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			x := next(cur.priority)
			if x < dist[v][0] {
				// first time reach v
				if dist[v][0] < math.MaxInt32 {
					heap.Push(&pq, NewItem(v*2+1, dist[v][0]))
				}
				dist[v][1] = dist[v][0]
				dist[v][0] = x
				heap.Push(&pq, NewItem(v*2, x))
			} else if x != dist[v][0] && x < dist[v][1] {
				// second time reach v
				dist[v][1] = x
				heap.Push(&pq, NewItem(v*2+1, x))
			}
		}
	}

	return dist[n-1][1]
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

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem(v int, p int) *Item {
	item := new(Item)
	item.value = v
	item.priority = p
	return item
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
