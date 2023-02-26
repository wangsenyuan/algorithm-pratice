package p2577

import "container/heap"

const INF = 1 << 30

var dd = []int{-1, 0, 1, 0, -1}

func minimumTime(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	pq := make(PriorityQueue, m*n)

	items := make([][]*Item, m)
	// always has a solution
	for i := 0; i < m; i++ {
		items[i] = make([]*Item, n)
		for j := 0; j < n; j++ {
			it := new(Item)
			it.x = i
			it.y = j
			it.priority = INF
			it.index = i*n + j
			items[i][j] = it
			pq[i*n+j] = it
		}
	}

	heap.Init(&pq)

	pq.update(items[0][0], 0)

	if grid[0][1] <= 1 {
		pq.update(items[0][1], 1)
	}
	if grid[1][0] <= 1 {
		pq.update(items[1][0], 1)
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		x, y := it.x, it.y

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && items[u][v].index >= 0 {
				di := grid[u][v] - it.priority
				di = max(di/2*2, 0) + 1
				if it.priority+di < items[u][v].priority {
					pq.update(items[u][v], it.priority+di)
				}
			}
		}
	}

	return items[m-1][n-1].priority
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	x        int
	y        int
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

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
