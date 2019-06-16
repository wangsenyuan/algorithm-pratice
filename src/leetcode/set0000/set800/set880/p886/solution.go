package p886

import (
	"container/heap"
	"math"
)

func reachableNodes(edges [][]int, M int, N int) int {
	conn := make([]map[int]int, N)
	for i := 0; i < N; i++ {
		conn[i] = make(map[int]int)
	}

	for _, edge := range edges {
		u, v, n := edge[0], edge[1], edge[2]
		conn[u][v] = n + 1
		conn[v][u] = n + 1
	}

	nodes := make([]*Item, N)
	for i := 0; i < N; i++ {
		item := new(Item)
		item.dist = math.MaxInt32
		item.value = i
		nodes[i] = item
	}
	nodes[0].dist = 0

	que := make(PriorityQueue, 0, N)

	for i := 0; i < N; i++ {
		heap.Push(&que, nodes[i])
	}

	for que.Len() > 0 {
		cur := heap.Pop(&que).(*Item)
		u := cur.value
		for v, d := range conn[u] {
			if nodes[v].dist > cur.dist+d {
				que.update(nodes[v], cur.dist+d)
			}
		}
	}

	var res int

	for i := 0; i < N; i++ {
		if nodes[i].dist <= M {
			res++
		}
	}

	for _, edge := range edges {
		u, v, d := edge[0], edge[1], edge[2]
		if nodes[u].dist <= M {
			x := M - nodes[u].dist
			c := min(d, x)
			res += c
			d -= c
		}
		if nodes[v].dist <= M {
			x := M - nodes[v].dist
			c := min(d, x)
			res += c
			d -= c
		}
	}

	return res
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
	dist  int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].dist < pq[j].dist
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, dist int) {
	item.dist = dist
	heap.Fix(pq, item.index)
}
