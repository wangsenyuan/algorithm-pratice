package p3386

import "container/heap"

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {

	g1 := make(map[string]map[string]float64)

	add := func(g map[string]map[string]float64, a string, b string, r float64) {
		if g[a] == nil {
			g[a] = make(map[string]float64)
		}
		g[a][b] = r
	}

	for i, cur := range pairs1 {
		a, b := cur[0], cur[1]
		r := rates1[i]
		add(g1, a, b, r)
		add(g1, b, a, 1/r)
	}

	g2 := make(map[string]map[string]float64)

	for i, cur := range pairs2 {
		a, b := cur[0], cur[1]
		r := rates2[i]
		add(g2, a, b, r)
		add(g2, b, a, 1/r)
	}

	pq := make(PriorityQueue, 0, 10)
	heap.Push(&pq, NewItem(initialCurrency, 1))

	best := make(map[string]float64)
	best[initialCurrency] = 1.0

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id

		if it.priority < best[u] {
			continue
		}

		if vs, ok := g1[u]; ok {
			for k, v := range vs {
				if best[k] < v*it.priority {
					best[k] = v * it.priority
					heap.Push(&pq, NewItem(k, best[k]))
				}
			}
		}
	}

	for k, v := range best {
		heap.Push(&pq, NewItem(k, v))
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id

		if it.priority < best[u] {
			continue
		}

		if vs, ok := g2[u]; ok {
			for k, v := range vs {
				if best[k] < v*it.priority {
					best[k] = v * it.priority
					heap.Push(&pq, NewItem(k, best[k]))
				}
			}
		}
	}

	return best[initialCurrency]
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       string  // The value of the item; arbitrary.
	priority float64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem(id string, v float64) *Item {
	it := new(Item)
	it.id = id
	it.priority = v
	return it
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

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
