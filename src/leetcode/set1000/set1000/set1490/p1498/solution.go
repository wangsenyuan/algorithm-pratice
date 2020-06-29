package p1498

import (
	"container/heap"
	"math"
)

func findMaxValueOfEquation(points [][]int, k int) int {
	// yi + yj + xj - xi 最大 when i < j
	// yi - xi + yj + xj
	n := len(points)

	var best = math.MinInt32

	items := make([]*Item, n)
	pq := make(PriorityQueue, 0, n)

	for i, j := 0, 0; i < n; i++ {
		p := points[i]
		x, y := p[0], p[1]

		for j < i && x-items[j].pos > k {
			pq.update(items[j], math.MinInt32)
			j++
		}

		if pq.Len() > 0 && pq[0].priority > math.MinInt32 {
			cur := y + x + pq[0].priority
			best = max(best, cur)
		}

		item := new(Item)
		item.pos = x
		item.priority = y - x
		items[i] = item
		heap.Push(&pq, item)
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	pos      int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
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
