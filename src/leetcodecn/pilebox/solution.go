package pilebox

import "container/heap"

func pileBox(box [][]int) int {
	n := len(box)

	outs := make([][]int, n)
	degree := make([]int, n)

	canPlace := func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if box[i][k] <= box[j][k] {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 10)

		for j := 0; j < n; j++ {
			if canPlace(i, j) {
				degree[j]++
				outs[i] = append(outs[i], j)
			}
		}
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		items[i] = &Item{i, degree[i], i}
		pq[i] = items[i]
	}

	heap.Init(&pq)

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = box[i][2]
	}
	var best int
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		i := cur.value
		if res[i] > best {
			best = res[i]
		}

		for _, j := range outs[i] {
			degree[j]--
			res[j] = max(res[j], res[i]+box[j][2])
			pq.update(items[j], degree[j])
		}
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
	value  int // The value of the item; arbitrary.
	degree int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].degree < pq[j].degree
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
func (pq *PriorityQueue) update(item *Item, degree int) {
	item.degree = degree
	heap.Fix(pq, item.index)
}
